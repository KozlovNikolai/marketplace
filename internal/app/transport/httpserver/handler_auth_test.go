package httpserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/KozlovNikolai/marketplace/internal/app/domain"
	"github.com/KozlovNikolai/marketplace/internal/app/transport/httpserver/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSignUp(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUserService := mocks.NewIUserService(t)
	h := HttpServer{
		userService: mockUserService,
	}

	testCases := []struct {
		name                 string
		inHandler            UserRequest
		wantErr              bool
		mockCreateUser       func()
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "Invalid JSON",
			wantErr:              true,
			mockCreateUser:       func() {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"invalid-json":"EOF"}`,
		},
		{
			name: "Invalid login request validation",
			inHandler: UserRequest{
				Login:    "cmd@cmdru",
				Password: "1234567",
			},
			wantErr:              true,
			mockCreateUser:       func() {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: "validation",
		},
		{
			name: "Invalid password request validation",
			inHandler: UserRequest{
				Login:    "cmd@cmd.ru",
				Password: "12345",
			},
			wantErr:              true,
			mockCreateUser:       func() {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: "validation",
		},

		{
			name: "Invalid service user create",
			inHandler: UserRequest{
				Login:    "cmd@cmd.ru",
				Password: "123456",
			},
			wantErr: true,
			mockCreateUser: func() {
				mockUserService.On("CreateUser", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("domain.User")).
					Return(domain.User{},
						domain.ErrDbCreationFailed).Once()
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"error service User":"data base creation failed"}`,
		},
		{
			name: "OK",
			inHandler: UserRequest{
				Login:    "cmd@cmd.ru",
				Password: "1234567",
			},
			wantErr: false,
			mockCreateUser: func() {
				mockUserService.On("CreateUser", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("domain.User")).
					Return(domain.NewUser(domain.NewUserData{
						ID:        1,
						Login:     "cmd@cmd.ru",
						Password:  "hashpassword",
						Role:      "regular",
						Token:     "",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					}),
						nil).Once()
			},
			expectedStatusCode:   http.StatusCreated,
			expectedResponseBody: `{"id":1,"login":"cmd@cmd.ru","password":"hashpassword","role":"regular","token":""}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Создаем новый контекст gin для тестирования

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Устанавливаем JSON-данные в запрос
			body, _ := json.Marshal(tc.inHandler)
			if tc.expectedResponseBody == `{"invalid-json":"EOF"}` {
				body = nil
			}
			c.Request = httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
			c.Request.Header.Set("Content-Type", "application/json")
			tc.mockCreateUser()
			h.SignUp(c)

			if tc.wantErr {
				assert.Equal(t, tc.expectedStatusCode, w.Code)
				require.Contains(t, w.Body.String(), tc.expectedResponseBody)
			} else {
				require.Equal(t, tc.expectedStatusCode, w.Code)
				require.Equal(t, tc.expectedResponseBody, w.Body.String())
			}
			mockUserService.AssertExpectations(t)
		})

	}
}

func TestSignIn(t *testing.T) {

	password := "123456"
	hash, err := hashPassword(password)
	assert.NoError(t, err)

	gin.SetMode(gin.TestMode)
	mockUserService := mocks.NewIUserService(t)
	mockTokenService := mocks.NewITokenService(t)
	h := HttpServer{
		userService:  mockUserService,
		tokenService: mockTokenService,
	}

	testCases := []struct {
		name                 string
		inHandler            UserRequest
		wantErr              bool
		mockGetUserByLogin   func()
		mockGenerateToken    func()
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "Invalid JSON",
			wantErr:              true,
			mockGetUserByLogin:   func() {},
			mockGenerateToken:    func() {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"invalid-json":"EOF"}`,
		},
		{
			name: "Invalid login request validation",
			inHandler: UserRequest{
				Login:    "cmd@cmdru",
				Password: "1234567",
			},
			wantErr:              true,
			mockGetUserByLogin:   func() {},
			mockGenerateToken:    func() {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: "validation",
		},
		{
			name: "Invalid password request validation",
			inHandler: UserRequest{
				Login:    "cmd@cmd.ru",
				Password: "12345",
			},
			wantErr:              true,
			mockGetUserByLogin:   func() {},
			mockGenerateToken:    func() {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: "validation",
		},
		{
			name: "Invalid service get user by login",
			inHandler: UserRequest{
				Login:    "cmd@cmd.ru",
				Password: password,
			},
			wantErr: true,
			mockGetUserByLogin: func() {
				mockUserService.On("GetUserByLogin", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("string")).
					Return(domain.User{},
						domain.ErrFailure).Once()
			},
			mockGenerateToken:    func() {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"invalid-request":"failure"}`,
		},
		{
			name: "Invalid service get user by login",
			inHandler: UserRequest{
				Login:    "cmd@cmd.ru",
				Password: password,
			},
			wantErr: true,
			mockGetUserByLogin: func() {
				mockUserService.On("GetUserByLogin", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("string")).
					Return(domain.NewUser(domain.NewUserData{
						ID:        1,
						Login:     "cmd@cmd.ru",
						Password:  "fake hash",
						Role:      "regular",
						Token:     "",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					}),
						nil).Once()
			},
			mockGenerateToken:    func() {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"invalid-password"}`,
		},
		{
			name: "OK",
			inHandler: UserRequest{
				Login:    "cmd@cmd.ru",
				Password: password,
			},
			wantErr: false,
			mockGetUserByLogin: func() {
				mockUserService.On("GetUserByLogin", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("string")).
					Return(domain.NewUser(domain.NewUserData{
						ID:        1,
						Login:     "cmd@cmd.ru",
						Password:  hash,
						Role:      "regular",
						Token:     "",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					}),
						nil).Once()
			},
			mockGenerateToken: func() {
				mockTokenService.On("GenerateToken", mock.AnythingOfType("domain.User")).
					Return(
						"tokenexample",
						nil).Once()
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"token":"tokenexample"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Создаем новый контекст gin для тестирования

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Устанавливаем JSON-данные в запрос
			body, _ := json.Marshal(tc.inHandler)
			if tc.expectedResponseBody == `{"invalid-json":"EOF"}` {
				body = nil
			}
			c.Request = httptest.NewRequest(http.MethodPost, "/signin", bytes.NewBuffer(body))
			c.Request.Header.Set("Content-Type", "application/json")
			tc.mockGetUserByLogin()
			tc.mockGenerateToken()
			h.SignIn(c)

			if tc.wantErr {
				assert.Equal(t, tc.expectedStatusCode, w.Code)
				require.Contains(t, w.Body.String(), tc.expectedResponseBody)
			} else {
				require.Equal(t, tc.expectedStatusCode, w.Code)
				require.Equal(t, tc.expectedResponseBody, w.Body.String())
			}
			mockUserService.AssertExpectations(t)
		})

	}
}

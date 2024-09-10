package httpserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KozlovNikolai/marketplace/internal/app/domain"
	"github.com/KozlovNikolai/marketplace/internal/app/transport/httpserver/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateProvider(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockProviderService := new(mocks.MockProviderService)

	h := HttpServer{
		providerService: mockProviderService,
	}

	testCases := []struct {
		name                 string
		inHandler            ProviderRequest
		wantErr              bool
		mockCreateProvider   func()
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "Invalid JSON",
			wantErr:              true,
			mockCreateProvider:   func() {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"invalid-json":"EOF"}`,
		},
		{
			name: "Invalid request validation",
			inHandler: ProviderRequest{
				Name:   "",
				Origin: "",
			},
			wantErr:              true,
			mockCreateProvider:   func() {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: "validation",
		},
		{
			name: "Invalid service provider response",
			inHandler: ProviderRequest{
				Name:   "Nissan",
				Origin: "Japan",
			},
			wantErr: true,
			mockCreateProvider: func() {
				mockProviderService.On("CreateProvider", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("domain.Provider")).
					Return(
						domain.NewProvider(domain.NewProviderData{ID: 1, Name: "TName", Origin: "TOrigin"}),
						domain.ErrDbCreationFailed).Once()
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"error service provider":"data base creation failed"}`,
		},
		{
			name: "OK",
			inHandler: ProviderRequest{
				Name:   "Nissan",
				Origin: "Japan",
			},
			wantErr: false,
			mockCreateProvider: func() {
				mockProviderService.On("CreateProvider", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("domain.Provider")).
					Return(domain.NewProvider(domain.NewProviderData{
						ID:     1,
						Name:   "Toyota",
						Origin: "Japan",
					}),
						nil).Once()
			},
			expectedStatusCode:   http.StatusCreated,
			expectedResponseBody: `{"id":1,"name":"Toyota","origin":"Japan"}`,
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
			c.Request = httptest.NewRequest(http.MethodPost, "/provider", bytes.NewBuffer(body))
			c.Request.Header.Set("Content-Type", "application/json")
			tc.mockCreateProvider()
			h.CreateProvider(c)

			if tc.wantErr {
				assert.Equal(t, tc.expectedStatusCode, w.Code)
				require.Contains(t, w.Body.String(), tc.expectedResponseBody)
			} else {
				require.Equal(t, tc.expectedStatusCode, w.Code)
				require.Equal(t, tc.expectedResponseBody, w.Body.String())
			}
			mockProviderService.AssertExpectations(t)
		})

	}
}

func TestGetProvider(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockProviderService := new(mocks.MockProviderService)
	h := HttpServer{
		providerService: mockProviderService,
	}
	testCases := []struct {
		name               string
		inHandler          string
		wantErr            bool
		mockGetProvider    func()
		expectStatusCode   int
		expectResponseBody string
	}{
		{
			name:      "OK",
			inHandler: "?id=3",
			wantErr:   false,
			mockGetProvider: func() {
				mockProviderService.On("GetProvider", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("int")).
					Return(domain.NewProvider(domain.NewProviderData{
						ID:     3,
						Name:   "AZLK",
						Origin: "USSR",
					}), nil).Once()
			},
			expectStatusCode:   http.StatusOK,
			expectResponseBody: `{"id":3,"name":"AZLK","origin":"USSR"}`,
		},
		{
			name:               "error validation id",
			inHandler:          "",
			wantErr:            true,
			mockGetProvider:    func() {},
			expectStatusCode:   http.StatusBadRequest,
			expectResponseBody: `invalid-provider-id`,
		},
		{
			name:               "invalid id=0",
			inHandler:          "?id=0",
			wantErr:            true,
			mockGetProvider:    func() {},
			expectStatusCode:   http.StatusBadRequest,
			expectResponseBody: `{"error":"id lower or equal zero"}`,
		},
		{
			name:               "invalid id<0",
			inHandler:          "?id=-4",
			wantErr:            true,
			mockGetProvider:    func() {},
			expectStatusCode:   http.StatusBadRequest,
			expectResponseBody: `{"error":"id lower or equal zero"}`,
		},
		{
			name:      "provider not found",
			inHandler: "?id=3",
			wantErr:   true,
			mockGetProvider: func() {
				mockProviderService.On("GetProvider", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("int")).
					Return(domain.NewProvider(domain.NewProviderData{
						ID:     0,
						Name:   "",
						Origin: "",
					}),
						domain.ErrNotFound).Once()
			},
			expectStatusCode:   http.StatusNotFound,
			expectResponseBody: `{"provider":"not found"}`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			url := fmt.Sprintf("/provider%s", tc.inHandler)
			c.Request = httptest.NewRequest(http.MethodGet, url, nil)
			c.Request.Header.Set("Content-Type", "application/json")
			tc.mockGetProvider()
			h.GetProvider(c)

			if tc.wantErr {
				require.Equal(t, tc.expectStatusCode, w.Code)
				require.Contains(t, w.Body.String(), tc.expectResponseBody)
			} else {
				require.Equal(t, tc.expectStatusCode, w.Code)
				require.Equal(t, tc.expectResponseBody, w.Body.String())
			}
		})
	}
}

func TestGetProviders(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockProviderService := new(mocks.MockProviderService)
	h := HttpServer{
		providerService: mockProviderService,
	}
	testCases := []struct {
		name               string
		inHandlerLimit     string
		inHandlerOffset    string
		wantErr            bool
		mockGetProvider    func()
		expectStatusCode   int
		expectResponseBody string
	}{
		{
			name:            "OK",
			inHandlerLimit:  "limit=10",
			inHandlerOffset: "offset=0",
			wantErr:         false,
			mockGetProvider: func() {
				mockProviderService.On("GetProviders", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
					Return([]domain.Provider{
						domain.NewProvider(domain.NewProviderData{ID: 1, Name: "AZLK", Origin: "USSR"}),
						domain.NewProvider(domain.NewProviderData{ID: 2, Name: "VAZ", Origin: "USSR"}),
						domain.NewProvider(domain.NewProviderData{ID: 3, Name: "GAZ", Origin: "USSR"}),
					}, nil).Once()
			},
			expectStatusCode:   http.StatusOK,
			expectResponseBody: `[{"id":1,"name":"AZLK","origin":"USSR"},{"id":2,"name":"VAZ","origin":"USSR"},{"id":3,"name":"GAZ","origin":"USSR"}]`,
		},
		{
			name:               "error get query limit",
			inHandlerLimit:     "",
			inHandlerOffset:    "offset=0",
			wantErr:            true,
			mockGetProvider:    func() {},
			expectStatusCode:   http.StatusBadRequest,
			expectResponseBody: `limit`,
		},
		{
			name:               "error get query offset",
			inHandlerLimit:     "limit=10",
			inHandlerOffset:    "",
			wantErr:            true,
			mockGetProvider:    func() {},
			expectStatusCode:   http.StatusBadRequest,
			expectResponseBody: `offset`,
		},
		{
			name:               "invalid limit",
			inHandlerLimit:     "limit=0",
			inHandlerOffset:    "offset=0",
			wantErr:            true,
			mockGetProvider:    func() {},
			expectStatusCode:   http.StatusBadRequest,
			expectResponseBody: `limit-must-be-greater-then-zero`,
		},
		{
			name:               "invalid offset",
			inHandlerLimit:     "limit=10",
			inHandlerOffset:    "offset=-4",
			wantErr:            true,
			mockGetProvider:    func() {},
			expectStatusCode:   http.StatusBadRequest,
			expectResponseBody: `offset-must-be-greater-or-equal-then-zero`,
		},
		{
			name:            "error service",
			inHandlerLimit:  "limit=10",
			inHandlerOffset: "offset=0",
			wantErr:         true,
			mockGetProvider: func() {
				mockProviderService.On("GetProviders", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
					Return([]domain.Provider{domain.NewProvider(domain.NewProviderData{ID: 0, Name: "", Origin: ""})}, domain.ErrNotFound).Once()
			},
			expectStatusCode:   http.StatusInternalServerError,
			expectResponseBody: `"error get providers":"not found"`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			url := fmt.Sprintf("/providers?%s&%s", tc.inHandlerLimit, tc.inHandlerOffset)
			c.Request = httptest.NewRequest(http.MethodGet, url, nil)
			c.Request.Header.Set("Content-Type", "application/json")
			tc.mockGetProvider()
			h.GetProviders(c)

			if tc.wantErr {
				require.Equal(t, tc.expectStatusCode, w.Code)
				require.Contains(t, w.Body.String(), tc.expectResponseBody)
			} else {
				require.Equal(t, tc.expectStatusCode, w.Code)
				require.Equal(t, tc.expectResponseBody, w.Body.String())
			}
		})
	}
}

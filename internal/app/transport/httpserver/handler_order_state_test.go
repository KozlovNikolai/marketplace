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

func TestCreateOrderState(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockOrderStateService := mocks.NewIOrderStateService(t)

	h := HttpServer{
		orderStateService: mockOrderStateService,
	}

	testCases := []struct {
		name                 string
		inHandler            OrderStateRequest
		wantErr              bool
		mockCreateOrderState func()
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "Invalid JSON",
			wantErr:              true,
			mockCreateOrderState: func() {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"invalid-json":"EOF"}`,
		},
		{
			name: "Invalid request validation",
			inHandler: OrderStateRequest{
				Name: "",
			},
			wantErr:              true,
			mockCreateOrderState: func() {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: "validation",
		},
		{
			name: "Invalid service OrderState response",
			inHandler: OrderStateRequest{
				Name: "Created",
			},
			wantErr: true,
			mockCreateOrderState: func() {
				mockOrderStateService.On("CreateOrderState", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("domain.OrderState")).
					Return(
						domain.NewOrderState(domain.NewOrderStateData{ID: 1, Name: "Created"}),
						domain.ErrDbCreationFailed).Once()
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"error service OrderState":"data base creation failed"}`,
		},
		{
			name: "OK",
			inHandler: OrderStateRequest{
				Name: "In progress",
			},
			wantErr: false,
			mockCreateOrderState: func() {
				mockOrderStateService.On("CreateOrderState", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("domain.OrderState")).
					Return(domain.NewOrderState(domain.NewOrderStateData{
						ID:   1,
						Name: "In progress",
					}),
						nil).Once()
			},
			expectedStatusCode:   http.StatusCreated,
			expectedResponseBody: `{"id":1,"name":"In progress"}`,
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
			c.Request = httptest.NewRequest(http.MethodPost, "/OrderState", bytes.NewBuffer(body))
			c.Request.Header.Set("Content-Type", "application/json")
			tc.mockCreateOrderState()
			h.CreateOrderState(c)

			if tc.wantErr {
				assert.Equal(t, tc.expectedStatusCode, w.Code)
				require.Contains(t, w.Body.String(), tc.expectedResponseBody)
			} else {
				require.Equal(t, tc.expectedStatusCode, w.Code)
				require.Equal(t, tc.expectedResponseBody, w.Body.String())
			}
			mockOrderStateService.AssertExpectations(t)
		})

	}
}

func TestGetOrderState(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockOrderStateService := mocks.NewIOrderStateService(t)
	h := HttpServer{
		orderStateService: mockOrderStateService,
	}
	testCases := []struct {
		name               string
		inHandler          string
		wantErr            bool
		mockGetOrderState  func()
		expectStatusCode   int
		expectResponseBody string
	}{
		{
			name:      "OK",
			inHandler: "?id=3",
			wantErr:   false,
			mockGetOrderState: func() {
				mockOrderStateService.On("GetOrderState", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("int")).
					Return(domain.NewOrderState(domain.NewOrderStateData{
						ID:   3,
						Name: "In delivery",
					}), nil).Once()
			},
			expectStatusCode:   http.StatusOK,
			expectResponseBody: `{"id":3,"name":"In delivery"}`,
		},
		{
			name:               "error validation id",
			inHandler:          "",
			wantErr:            true,
			mockGetOrderState:  func() {},
			expectStatusCode:   http.StatusBadRequest,
			expectResponseBody: `invalid-orderState-id`,
		},
		{
			name:               "invalid id=0",
			inHandler:          "?id=0",
			wantErr:            true,
			mockGetOrderState:  func() {},
			expectStatusCode:   http.StatusBadRequest,
			expectResponseBody: `{"error":"id lower or equal zero"}`,
		},
		{
			name:               "invalid id lower then zero",
			inHandler:          "?id=-4",
			wantErr:            true,
			mockGetOrderState:  func() {},
			expectStatusCode:   http.StatusBadRequest,
			expectResponseBody: `{"error":"id lower or equal zero"}`,
		},
		{
			name:      "orderState not found",
			inHandler: "?id=3",
			wantErr:   true,
			mockGetOrderState: func() {
				mockOrderStateService.On("GetOrderState", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("int")).
					Return(domain.NewOrderState(domain.NewOrderStateData{
						ID:   0,
						Name: "",
					}),
						domain.ErrNotFound).Once()
			},
			expectStatusCode:   http.StatusNotFound,
			expectResponseBody: `{"orderState":"not found"}`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			url := fmt.Sprintf("/OrderState%s", tc.inHandler)
			c.Request = httptest.NewRequest(http.MethodGet, url, nil)
			c.Request.Header.Set("Content-Type", "application/json")
			tc.mockGetOrderState()
			h.GetOrderState(c)

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

func TestGetOrderStates(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockOrderStateService := mocks.NewIOrderStateService(t)
	h := HttpServer{
		orderStateService: mockOrderStateService,
	}
	testCases := []struct {
		name               string
		inHandlerLimit     string
		inHandlerOffset    string
		wantErr            bool
		mockGetOrderState  func()
		expectStatusCode   int
		expectResponseBody string
	}{
		{
			name:            "OK",
			inHandlerLimit:  "limit=10",
			inHandlerOffset: "offset=0",
			wantErr:         false,
			mockGetOrderState: func() {
				mockOrderStateService.On("GetOrderStates", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
					Return([]domain.OrderState{
						domain.NewOrderState(domain.NewOrderStateData{ID: 1, Name: "Created"}),
						domain.NewOrderState(domain.NewOrderStateData{ID: 2, Name: "In progress"}),
						domain.NewOrderState(domain.NewOrderStateData{ID: 3, Name: "In delivery"}),
					}, nil).Once()
			},
			expectStatusCode:   http.StatusOK,
			expectResponseBody: `[{"id":1,"name":"Created"},{"id":2,"name":"In progress"},{"id":3,"name":"In delivery"}]`,
		},
		{
			name:               "error get query limit",
			inHandlerLimit:     "",
			inHandlerOffset:    "offset=0",
			wantErr:            true,
			mockGetOrderState:  func() {},
			expectStatusCode:   http.StatusBadRequest,
			expectResponseBody: `limit`,
		},
		{
			name:               "error get query offset",
			inHandlerLimit:     "limit=10",
			inHandlerOffset:    "",
			wantErr:            true,
			mockGetOrderState:  func() {},
			expectStatusCode:   http.StatusBadRequest,
			expectResponseBody: `offset`,
		},
		{
			name:               "invalid limit",
			inHandlerLimit:     "limit=0",
			inHandlerOffset:    "offset=0",
			wantErr:            true,
			mockGetOrderState:  func() {},
			expectStatusCode:   http.StatusBadRequest,
			expectResponseBody: `limit-must-be-greater-then-zero`,
		},
		{
			name:               "invalid offset",
			inHandlerLimit:     "limit=10",
			inHandlerOffset:    "offset=-4",
			wantErr:            true,
			mockGetOrderState:  func() {},
			expectStatusCode:   http.StatusBadRequest,
			expectResponseBody: `offset-must-be-greater-or-equal-then-zero`,
		},
		{
			name:            "error service",
			inHandlerLimit:  "limit=10",
			inHandlerOffset: "offset=0",
			wantErr:         true,
			mockGetOrderState: func() {
				mockOrderStateService.On("GetOrderStates", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
					Return([]domain.OrderState{domain.NewOrderState(domain.NewOrderStateData{ID: 0, Name: ""})}, domain.ErrNotFound).Once()
			},
			expectStatusCode:   http.StatusInternalServerError,
			expectResponseBody: `"error get orderStates":"not found"`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			url := fmt.Sprintf("/OrderStates?%s&%s", tc.inHandlerLimit, tc.inHandlerOffset)
			c.Request = httptest.NewRequest(http.MethodGet, url, nil)
			c.Request.Header.Set("Content-Type", "application/json")
			tc.mockGetOrderState()
			h.GetOrderStates(c)

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

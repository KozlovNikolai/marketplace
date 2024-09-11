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

func TestCreateProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockProductService := mocks.NewIProductService(t)
	h := HttpServer{
		productService: mockProductService,
	}

	testCases := []struct {
		name                 string
		inHandler            ProductRequest
		wantErr              bool
		mockCreateProduct    func()
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "Invalid JSON",
			wantErr:              true,
			mockCreateProduct:    func() {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"invalid-json":"EOF"}`,
		},
		{
			name: "Invalid name request validation",
			inHandler: ProductRequest{
				Name:       "",
				ProviderID: 1,
				Price:      123.45,
				Stock:      34,
			},
			wantErr:              true,
			mockCreateProduct:    func() {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: "validation",
		},
		{
			name: "Invalid Product id request validation",
			inHandler: ProductRequest{
				Name:       "AZLK-2140",
				ProviderID: 0,
				Price:      123.45,
				Stock:      34,
			},
			wantErr:              true,
			mockCreateProduct:    func() {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: "validation",
		},
		{
			name: "Invalid price request validation",
			inHandler: ProductRequest{
				Name:       "AZLK-2140",
				ProviderID: 1,
				Price:      -4,
				Stock:      34,
			},
			wantErr:              true,
			mockCreateProduct:    func() {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: "validation",
		},
		{
			name: "Invalid stock request validation",
			inHandler: ProductRequest{
				Name:       "AZLK-2140",
				ProviderID: 1,
				Price:      123.45,
				Stock:      -6,
			},
			wantErr:              true,
			mockCreateProduct:    func() {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: "validation",
		},
		{
			name: "service error",
			inHandler: ProductRequest{
				Name:       "AZLK-2140",
				ProviderID: 1,
				Price:      123.45,
				Stock:      34,
			},
			wantErr: true,
			mockCreateProduct: func() {
				mockProductService.On("CreateProduct", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("domain.Product")).
					Return(domain.Product{},
						domain.ErrFailure).Once()
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"error DB saving product":"failure"}`,
		},
		{
			name: "OK",
			inHandler: ProductRequest{
				Name:       "AZLK-2140",
				ProviderID: 1,
				Price:      123.45,
				Stock:      34,
			},
			wantErr: false,
			mockCreateProduct: func() {
				mockProductService.On("CreateProduct", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("domain.Product")).
					Return(domain.NewProduct(domain.NewProductData{
						ID:         1,
						Name:       "AZLK-2140",
						ProviderID: 1,
						Price:      123.45,
						Stock:      34,
					}),
						nil).Once()
			},
			expectedStatusCode:   http.StatusCreated,
			expectedResponseBody: `{"id":1,"name":"AZLK-2140","provider_id":1,"price":123.45,"stock":34}`,
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
			c.Request = httptest.NewRequest(http.MethodPost, "/Product", bytes.NewBuffer(body))
			c.Request.Header.Set("Content-Type", "application/json")
			tc.mockCreateProduct()
			h.CreateProduct(c)

			if tc.wantErr {
				assert.Equal(t, tc.expectedStatusCode, w.Code)
				require.Contains(t, w.Body.String(), tc.expectedResponseBody)
			} else {
				require.Equal(t, tc.expectedStatusCode, w.Code)
				require.Equal(t, tc.expectedResponseBody, w.Body.String())
			}
			mockProductService.AssertExpectations(t)
		})

	}
}

func TestGetProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockProductService := mocks.NewIProductService(t)
	h := HttpServer{
		productService: mockProductService,
	}
	testCases := []struct {
		name               string
		inHandler          string
		wantErr            bool
		mockGetProduct     func()
		expectStatusCode   int
		expectResponseBody string
	}{
		{
			name:      "OK",
			inHandler: "?id=1",
			wantErr:   false,
			mockGetProduct: func() {
				mockProductService.On("GetProduct", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("int")).
					Return(domain.NewProduct(domain.NewProductData{
						ID:         1,
						Name:       "AZLK-2140",
						ProviderID: 1,
						Price:      123.45,
						Stock:      34,
					}), nil).Once()
			},
			expectStatusCode:   http.StatusOK,
			expectResponseBody: `{"id":1,"name":"AZLK-2140","provider_id":1,"price":123.45,"stock":34}`,
		},
		{
			name:               "error validation id",
			inHandler:          "",
			wantErr:            true,
			mockGetProduct:     func() {},
			expectStatusCode:   http.StatusBadRequest,
			expectResponseBody: `invalid-product-id`,
		},
		{
			name:               "invalid id lower or equal zero",
			inHandler:          "?id=0",
			wantErr:            true,
			mockGetProduct:     func() {},
			expectStatusCode:   http.StatusBadRequest,
			expectResponseBody: `{"error":"id lower or equal zero"}`,
		},
		{
			name:      "Product not found",
			inHandler: "?id=3",
			wantErr:   true,
			mockGetProduct: func() {
				mockProductService.On("GetProduct", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("int")).
					Return(domain.Product{},
						domain.ErrFailure).Once()
			},
			expectStatusCode:   http.StatusInternalServerError,
			expectResponseBody: `{"error-get-product":"failure"}`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			url := fmt.Sprintf("/Product%s", tc.inHandler)
			c.Request = httptest.NewRequest(http.MethodGet, url, nil)
			c.Request.Header.Set("Content-Type", "application/json")
			tc.mockGetProduct()
			h.GetProduct(c)

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

func TestGetProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockProductService := mocks.NewIProductService(t)
	h := HttpServer{
		productService: mockProductService,
	}
	testCases := []struct {
		name               string
		inHandlerURL       string
		wantErr            bool
		mockGetProduct     func()
		expectStatusCode   int
		expectResponseBody string
	}{
		{
			name:         "OK",
			inHandlerURL: "?limit=10&offset=0",
			wantErr:      false,
			mockGetProduct: func() {
				mockProductService.On("GetProducts", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
					Return([]domain.Product{
						domain.NewProduct(domain.NewProductData{ID: 1, Name: "AZLK-2140", ProviderID: 1, Price: 123.45, Stock: 34}),
						domain.NewProduct(domain.NewProductData{ID: 2, Name: "VAZ", ProviderID: 2, Price: 1223.45, Stock: 3}),
						domain.NewProduct(domain.NewProductData{ID: 3, Name: "VOLGA", ProviderID: 21, Price: 1243.45, Stock: 346}),
					}, nil).Once()
			},
			expectStatusCode:   http.StatusOK,
			expectResponseBody: `[{"id":1,"name":"AZLK-2140","provider_id":1,"price":123.45,"stock":34},{"id":2,"name":"VAZ","provider_id":2,"price":1223.45,"stock":3},{"id":3,"name":"VOLGA","provider_id":21,"price":1243.45,"stock":346}]`,
		},
		{
			name:               "error get query limit",
			inHandlerURL:       "?offset=0",
			wantErr:            true,
			mockGetProduct:     func() {},
			expectStatusCode:   http.StatusBadRequest,
			expectResponseBody: `limit`,
		},
		{
			name:               "error get query offset",
			inHandlerURL:       "?limit=10",
			wantErr:            true,
			mockGetProduct:     func() {},
			expectStatusCode:   http.StatusBadRequest,
			expectResponseBody: `offset`,
		},
		{
			name:               "invalid limit",
			inHandlerURL:       "?limit=0&offset=0",
			wantErr:            true,
			mockGetProduct:     func() {},
			expectStatusCode:   http.StatusBadRequest,
			expectResponseBody: `limit-must-be-greater-then-zero`,
		},
		{
			name:               "invalid offset",
			inHandlerURL:       "?limit=10&offset=-4",
			wantErr:            true,
			mockGetProduct:     func() {},
			expectStatusCode:   http.StatusBadRequest,
			expectResponseBody: `offset-must-be-greater-or-equal-then-zero`,
		},
		{
			name:         "error service",
			inHandlerURL: "?limit=10&offset=0",
			wantErr:      true,
			mockGetProduct: func() {
				mockProductService.On("GetProducts", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).
					Return([]domain.Product{},
						domain.ErrFailure).Once()
			},
			expectStatusCode:   http.StatusInternalServerError,
			expectResponseBody: `"error get products":"failure"`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			url := fmt.Sprintf("/Products%s", tc.inHandlerURL)
			c.Request = httptest.NewRequest(http.MethodGet, url, nil)
			c.Request.Header.Set("Content-Type", "application/json")
			tc.mockGetProduct()
			h.GetProducts(c)

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

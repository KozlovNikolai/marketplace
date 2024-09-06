package httpserver

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestProductRequestValidate(t *testing.T) {
	testCases := []struct {
		name          string
		in            ProductRequest
		wantErr       bool
		expectedField []string
		expectedTag   []string
	}{
		{
			name: "valid product request data",
			in: ProductRequest{
				Name:       "valenki",
				ProviderID: 3,
				Price:      123.45,
				Stock:      45,
			},
			wantErr:       false,
			expectedField: []string{},
			expectedTag:   []string{},
		},
		{
			name: "valid product request data with zero values",
			in: ProductRequest{
				Name:       "valenki",
				ProviderID: 3,
				Price:      0,
				Stock:      0,
			},
			wantErr:       false,
			expectedField: []string{},
			expectedTag:   []string{},
		},
		{
			name: "ivalid name in product request data",
			in: ProductRequest{
				Name:       "",
				ProviderID: 3,
				Price:      123.45,
				Stock:      14,
			},
			wantErr:       true,
			expectedField: []string{"Name"},
			expectedTag:   []string{"required"},
		},
		{
			name: "ivalid provider id in product request data",
			in: ProductRequest{
				Name:       "val",
				ProviderID: 0,
				Price:      123.45,
				Stock:      14,
			},
			wantErr:       true,
			expectedField: []string{"ProviderID"},
			expectedTag:   []string{"gt"},
		},
		{
			name: "ivalid price in product request data",
			in: ProductRequest{
				Name:       "val",
				ProviderID: 3,
				Price:      -4,
				Stock:      14,
			},
			wantErr:       true,
			expectedField: []string{"Price"},
			expectedTag:   []string{"gte"},
		},
		{
			name: "ivalid stock in product request data",
			in: ProductRequest{
				Name:       "val",
				ProviderID: 3,
				Price:      4,
				Stock:      -3,
			},
			wantErr:       true,
			expectedField: []string{"Stock"},
			expectedTag:   []string{"gte"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.in.Validate()
			if tc.wantErr {
				validationErrors, ok := err.(validator.ValidationErrors)
				assert.True(t, ok)
				if len(validationErrors) > 0 {
					for i := 0; i < len(validationErrors); i++ {
						assert.Equal(t, tc.expectedField[i], validationErrors[i].Field())
						assert.Equal(t, tc.expectedTag[i], validationErrors[i].Tag())
					}
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
func TestOrderRequestValidate(t *testing.T) {
	testCases := []struct {
		name          string
		in            OrderRequest
		wantErr       bool
		expectedField []string
		expectedTag   []string
	}{
		{
			name: "valid order create request",
			in: OrderRequest{
				UserID: 1,
			},
			wantErr:       false,
			expectedField: []string{},
			expectedTag:   []string{},
		},
		{
			name: "invalid user id in order create request",
			in: OrderRequest{
				UserID: 0,
			},
			wantErr:       true,
			expectedField: []string{"UserID"},
			expectedTag:   []string{"gt"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.in.Validate()
			if tc.wantErr {
				validationErrors, ok := err.(validator.ValidationErrors)
				assert.True(t, ok)
				if len(validationErrors) > 0 {
					for i := 0; i < len(validationErrors); i++ {
						assert.Equal(t, tc.expectedField[i], validationErrors[i].Field())
						assert.Equal(t, tc.expectedTag[i], validationErrors[i].Tag())
					}
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
func TestProviderRequestValidate(t *testing.T) {
	testCases := []struct {
		name          string
		in            ProviderRequest
		wantErr       bool
		expectedField []string
		expectedTag   []string
	}{
		{
			name: "valid provider",
			in: ProviderRequest{
				Name:   "Microsoft",
				Origin: "Russia",
			},
			wantErr:       false,
			expectedField: []string{},
			expectedTag:   []string{},
		},
		{
			name: "empty provider name",
			in: ProviderRequest{
				Name:   "",
				Origin: "Russia",
			},
			wantErr:       true,
			expectedField: []string{"Name"},
			expectedTag:   []string{"required"},
		},
		{
			name: "empty provider origin",
			in: ProviderRequest{
				Name:   "Microsoft",
				Origin: "",
			},
			wantErr:       true,
			expectedField: []string{"Origin"},
			expectedTag:   []string{"required"},
		},
		{
			name: "empty provider name and origin",
			in: ProviderRequest{
				Name:   "",
				Origin: "",
			},
			wantErr:       true,
			expectedField: []string{"Name", "Origin"},
			expectedTag:   []string{"required", "required"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.in.Validate()
			if tc.wantErr {
				validationErrors, ok := err.(validator.ValidationErrors)
				assert.True(t, ok)
				if len(validationErrors) > 0 {
					for i := 0; i < len(validationErrors); i++ {
						assert.Equal(t, tc.expectedField[i], validationErrors[i].Field())
						assert.Equal(t, tc.expectedTag[i], validationErrors[i].Tag())
					}
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestOrderStateRequestValidate(t *testing.T) {
	testCases := []struct {
		name          string
		in            OrderStateRequest
		wantErr       bool
		expectedField []string
		expextedTag   []string
	}{
		{
			name: "valid order state",
			in: OrderStateRequest{
				Name: "in progres",
			},
			wantErr:       false,
			expectedField: []string{},
			expextedTag:   []string{},
		},
		{
			name: "valid order state",
			in: OrderStateRequest{
				Name: "",
			},
			wantErr:       true,
			expectedField: []string{"Name"},
			expextedTag:   []string{"required"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.in.Validate()
			if tc.wantErr {
				validationErrors, ok := err.(validator.ValidationErrors)
				assert.True(t, ok)
				if len(validationErrors) > 0 {
					for i := 0; i < len(validationErrors); i++ {
						assert.Equal(t, tc.expectedField[i], validationErrors[i].Field())
						assert.Equal(t, tc.expextedTag[i], validationErrors[i].Tag())
					}
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUserRequestValidate(t *testing.T) {
	cases := []struct {
		name          string
		in            UserRequest
		wantErr       bool
		expectedField []string
		expectedTag   []string
	}{
		{
			name: "valid user data",
			in: UserRequest{
				Login:    "cmd@cmd.ru",
				Password: "12345678901234567890123456789012",
			},
			wantErr: false,
		},
		{
			name: "empty login",
			in: UserRequest{
				Login:    "",
				Password: "12345678901234567890123456789012",
			},
			wantErr:       true,
			expectedField: []string{"Login"},
			expectedTag:   []string{"required"},
		},
		{
			name: "wrong login format",
			in: UserRequest{
				Login:    "cmd@cmdru",
				Password: "12345678901234567890123456789012",
			},
			wantErr:       true,
			expectedField: []string{"Login"},
			expectedTag:   []string{"email"},
		},
		{
			name: "empty password",
			in: UserRequest{
				Login:    "cmd@cmd.ru",
				Password: "",
			},
			wantErr:       true,
			expectedField: []string{"Password"},
			expectedTag:   []string{"required"},
		},
		{
			name: "short password error",
			in: UserRequest{
				Login:    "cmd@cmd.ru",
				Password: "123",
			},
			wantErr:       true,
			expectedField: []string{"Password"},
			expectedTag:   []string{"min"},
		},
		{
			name: "long password error",
			in: UserRequest{
				Password: "123456789012345678901234567890123",
				Login:    "cmd@cmd.ru",
			},
			wantErr:       true,
			expectedField: []string{"Password"},
			expectedTag:   []string{"max"},
		},
		{
			name: "long password error and wrong login",
			in: UserRequest{
				Login:    "cmd@cmdru",
				Password: "123456789012345678901234567890123",
			},
			wantErr:       true,
			expectedField: []string{"Login", "Password"},
			expectedTag:   []string{"email", "max"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := c.in.Validate()
			if c.wantErr {
				assert.Error(t, err)
				// Проверяем конкретные ошибки валидации
				validationErrors, ok := err.(validator.ValidationErrors)
				assert.True(t, ok)

				// Проверяем поле и тип ошибки
				if len(validationErrors) > 0 {
					for i := 0; i < len(validationErrors); i++ {
						ve := validationErrors[i]
						assert.Equal(t, c.expectedField[i], ve.Field())
						assert.Equal(t, c.expectedTag[i], ve.Tag())
					}
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestItemRequestValidate(t *testing.T) {
	testCases := []struct {
		name          string
		in            ItemRequest
		wantErr       bool
		expectedField string
		expectedTag   string
	}{
		{
			name: "valid item request data",
			in: ItemRequest{
				ProductID: 4,
				Quantity:  6,
				OrderID:   3,
			},
			wantErr:       false,
			expectedField: "",
			expectedTag:   "",
		},
		{
			name: "invalid product id in item request data",
			in: ItemRequest{
				ProductID: 0,
				Quantity:  6,
				OrderID:   3,
			},
			wantErr:       true,
			expectedField: "ProductID",
			expectedTag:   "gt",
		},
		{
			name: "invalid quantity in item request data",
			in: ItemRequest{
				ProductID: 3,
				Quantity:  0,
				OrderID:   3,
			},
			wantErr:       true,
			expectedField: "Quantity",
			expectedTag:   "gt",
		},
		{
			name: "invalid order id in item request data",
			in: ItemRequest{
				ProductID: 5,
				Quantity:  6,
				OrderID:   0,
			},
			wantErr:       true,
			expectedField: "OrderID",
			expectedTag:   "gt",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.in.Validate()
			if tc.wantErr {
				validationErrors, ok := err.(validator.ValidationErrors)
				assert.True(t, ok)
				if len(validationErrors) > 0 {
					assert.Equal(t, tc.expectedField, validationErrors[0].Field())
					assert.Equal(t, tc.expectedTag, validationErrors[0].Tag())
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

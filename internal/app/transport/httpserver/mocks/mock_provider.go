// mocks/ProviderService.go
package mocks

import (
	"context"

	"github.com/KozlovNikolai/marketplace/internal/app/domain"
	"github.com/stretchr/testify/mock"
)

type MockProviderService struct {
	mock.Mock
}

func (m *MockProviderService) CreateProvider(ctx context.Context, provider domain.Provider) (domain.Provider, error) {
	args := m.Called(ctx, provider)
	return args.Get(0).(domain.Provider), args.Error(1)
}

func (m *MockProviderService) GetProvider(ctx context.Context, id int) (domain.Provider, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Provider), args.Error(1)
}

func (m *MockProviderService) GetProviders(ctx context.Context, limit int, offset int) ([]domain.Provider, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).([]domain.Provider), args.Error(1)
}

func (m *MockProviderService) UpdateProvider(ctx context.Context, provider domain.Provider) (domain.Provider, error) {
	args := m.Called(ctx, provider)
	return args.Get(0).(domain.Provider), args.Error(1)
}

func (m *MockProviderService) DeleteProvider(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

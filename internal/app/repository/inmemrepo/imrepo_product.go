package inmemrepo

import (
	"context"
	"fmt"
	"sort"
	"sync"

	"github.com/KozlovNikolai/marketplace/internal/app/domain"
	"github.com/KozlovNikolai/marketplace/internal/app/repository/models"
)

type ProductRepo struct {
	db    *inMemStore
	mutex sync.RWMutex
}

func NewProductRepo(db *inMemStore) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

// CreateProduct implements services.IProductRepository.
func (repo *ProductRepo) CreateProduct(_ context.Context, product domain.Product) (domain.Product, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	// проверяем, существует ли поставщик
	if _, exists := repo.db.providers[product.ProviderID()]; !exists {
		return domain.Product{}, fmt.Errorf("provider with id %d does not exist", product.ProviderID())
	}
	// мапим домен в модель
	dbProduct := domainToProduct(product)
	dbProduct.ID = repo.db.nextProductsID

	// инкрементируем счетчик записей
	repo.db.nextProductsID++
	// сохраняем
	repo.db.products[dbProduct.ID] = dbProduct
	// мапим модель в домен
	domainProduct, err := productToDomain(dbProduct)
	if err != nil {
		return domain.Product{}, fmt.Errorf("failed to create domain Product: %w", err)
	}
	return domainProduct, nil
}

// DeleteProduct implements services.IProductRepository.
func (repo *ProductRepo) DeleteProduct(_ context.Context, id int) error {
	if id == 0 {
		return fmt.Errorf("%w: id", domain.ErrRequired)
	}
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	_, exists := repo.db.products[id]
	if !exists {
		return fmt.Errorf("product with id %d - %w", id, domain.ErrNotFound)
	}
	delete(repo.db.products, id)
	return nil
}

// GetProduct implements services.IProductRepository.
func (repo *ProductRepo) GetProduct(_ context.Context, id int) (domain.Product, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	Product, exists := repo.db.products[id]
	if !exists {
		return domain.Product{}, fmt.Errorf("product with id %d - %w", id, domain.ErrNotFound)
	}
	domainProduct, err := productToDomain(Product)
	if err != nil {
		return domain.Product{}, fmt.Errorf("failed to create domain Product: %w", err)
	}
	return domainProduct, nil
}

// GetProducts implements services.IProductRepository.
func (repo *ProductRepo) GetProducts(_ context.Context, limit int, offset int) ([]domain.Product, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	// извлекаем все ключи из мапы и сортируем их
	keys := make([]int, 0, len(repo.db.products))
	for k := range repo.db.products {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	// выбираем записи с нужными ключами
	var products []models.Product
	for i := offset; i < offset+limit && i < len(keys); i++ {
		products = append(products, repo.db.products[keys[i]])
	}

	// мапим массив моделей в массив доменов
	domainproducts := make([]domain.Product, len(products))
	for i, product := range products {
		domainproduct, err := productToDomain(product)
		if err != nil {
			return nil, fmt.Errorf("failed to create domain User: %w", err)
		}
		domainproducts[i] = domainproduct
	}
	return domainproducts, nil
}

// UpdateProduct implements services.IProductRepository.
func (repo *ProductRepo) UpdateProduct(_ context.Context, Product domain.Product) (domain.Product, error) {
	dbProduct := domainToProduct(Product)
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	// проверяем наличие записи
	_, exists := repo.db.products[dbProduct.ID]
	if !exists {
		return domain.Product{}, fmt.Errorf("product with id %d - %w", dbProduct.ID, domain.ErrNotFound)
	}
	// обновляем запись
	repo.db.products[dbProduct.ID] = dbProduct
	domainProduct, err := productToDomain(dbProduct)
	if err != nil {
		return domain.Product{}, fmt.Errorf("failed to create domain Product: %w", err)
	}
	return domainProduct, nil
}

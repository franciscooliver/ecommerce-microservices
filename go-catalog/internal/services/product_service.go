package services

import (
	"github.com/franciscooliver/imersao-fc/internal/database"
	"github.com/franciscooliver/imersao-fc/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(productDB *database.ProductDB) *ProductService {
	return &ProductService{ProductDB: *productDB}
}

func (ps *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) FindById(id string) (*entity.Product, error) {
	product, err := ps.ProductDB.FindById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) FindByCategoryId(categoryID string) ([]*entity.Product, error) {
	products, err := ps.ProductDB.FindByCategoryId(categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) CreateProduct(
	name,
	description string,
	price float64,
	categoryID string,
	imageURL string,
) (*entity.Product, error) {
	product := entity.NewProduct(name, description, price, categoryID, imageURL)
	_, err := ps.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

package entity

import "github.com/google/uuid"


type ProductRepository interface {
  Create(product *Product) error
  ListAll() ([]*Product, error)
}


type Product struct {
  ID string
  Name string
  Price float64
}

func NewProduct(name string, price float64) *Product{
  return &Product{
    ID: uuid.New().string(),
    Name: name,
    Price: price
  }
}

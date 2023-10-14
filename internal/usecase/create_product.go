package usecase

type CreateProductInputDto struct {
  Name string `json:"name"`
  Price float64  `json:"price"`
}

type CreateProductOutputDto struct {
  ID string
  Name string
  Price float64
}



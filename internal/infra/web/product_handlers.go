package web

import (
  "encoding/json"
  "net/http"
  "github.com/angelino-valeta/hexagonal-akakfa-golang-crud-product/internal/usecase"
)

type ProductHandlers struct {
  CreateProductUseCase *usecase.CreateProductUseCase
  ListProductUseCase *usecase.ListProductUseCase
}

func NewProductHandlers(createProductUseCase *usecase.CreateProductUseCase, listProductUseCase *usecase.ListProductUseCase) *ProductHandlers {
  return &ProductHandlers {
    CreateProductUseCase: createProductUseCase,
    ListProductUseCase: listProductUseCase
  }
}

func (p *ProductHandlers) CreateProductHandler(w http.ResponseWriter, r *http.Request){
  var input usecase.CreateProductInputDto
  err := json.newDecoder(r.Body).Decode(&input)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  output, err := p.CreateUseCase.Execute(input)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  json.NewEnconder(w).Enconde(output)
}

func (p *ProductHandlers) LstProductHandler(w http.ResponseWriter, r http.Request){
  output, err := p.ListProductUseCase.Execute()
  if err != nil {
    w.writeHeader(http.StatusInternalServerError)
    return
  }
  w.header().Set("Content-type", "application/json")
  w.WriteHeader(http.statusOk)
  json.NewEnconder(w).Enconde(output)
}

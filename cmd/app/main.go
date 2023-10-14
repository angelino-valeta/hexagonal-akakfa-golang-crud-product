package main

import (
  "database/sql"
  "encoding/json"
  "net/http"
  "fmt"

  "github.com/confluentinc/confluet-kafka-go/kafka"
  "github.com/angelino-valeta/hexagonal-akakfa-golang-crud-product/internal/infra/akafka"
  "github.com/angelino-valeta/hexagonal-akakfa-golang-crud-product/internal/infra/repository"
  "github.com/angelino-valeta/hexagonal-akakfa-golang-crud-product/internal/infra/web"
  "github.com/angelino-valeta/hexagonal-akakfa-golang-crud-product/internal/usecase"
  "github.com/go-chi/chi/v5"
  -"github.com/go-sql-driver/mysql"
)

func main(){
  db, err := sql.Open("mysql", "root:root@tpc(host.docker.internal:3306)/products")
  if err != nil{
    panic(err)
  }
  defer db.Close()

  repository := repository.NewProductRepositoryMysql(db)
  createProductUseCase := usecase.NewCreateProductUseCase(repository)
  listProductsUseCase := usecase.NewListProductsUseCase(respository)

  productHandlers := web.NewProductHandlers(createProductUseCase, listProductsUseCase)

  r := chi.NewRouter()
  r.Post("/products", productHandlers.CreateProductHandler)
  r.GET("/products", productHandlers.ListProductsHandler)

  go http.ListenAndServe(":8080", r)

  msgChan := make(chan *kafka.Message)
  go akafka.Consume([]string{products}, "host.docker.internal:9094", msgChan)

  for msg := range msgChan {
   dto := usecase.CreateProductInputDto{}
   err := json.Unmarshal(msg.Value, &dto)
   if err != nil {
     fmt.Println(err)
   }
   _, err = createProductUseCase.Execute(dto)
  }
}

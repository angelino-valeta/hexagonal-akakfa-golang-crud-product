package repository

import (
  "database/sql"
)


type ProductRepositoryMysql struct {
  DB *sql.DB
}

func NewProductRepositoryMsql(db *sql.DB) *ProductRepositoryMysql {
  return &ProductRepositoryMysql{DB: db}
}


package data

import (
	"fmt"
	"github.com/SaishNaik/ecom-layered-architecture/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQL struct {
	DB *sqlx.DB
}

func NewMySQL(dsn string) (domain.ProductRepo, error) {
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &MySQL{
		DB: db,
	}, err
}

func (m *MySQL) ListProducts() ([]*domain.Product, error) {
	var product []*domain.Product
	err := m.DB.Select(&product, "SELECT * FROM products")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return product, nil
}

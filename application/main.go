package main

import (
	"github.com/SaishNaik/ecom-layered-architecture/data"
	"github.com/SaishNaik/ecom-layered-architecture/domain"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

var repo domain.ProductRepo

func main() {
	//dsn := os.Getenv("mysql_dsn")
	dsn := "root:root@(localhost:3306)/test"
	var err error
	repo, err = data.NewMySQL(dsn)
	if err != nil {
		log.Panic(err)
	}

	e := echo.New()
	e.GET("/", getProducts)
	e.Logger.Fatal(e.Start(":1323"))
}

func getProducts(c echo.Context) error {
	products, err := repo.ListProducts()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, products)
}

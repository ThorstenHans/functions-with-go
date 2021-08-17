package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var products map[int]Product = map[int]Product{
	1: {"Milk"},
	2: {"Butter"},
}

type Product struct {
	Name string
}

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

func getProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid product identifier")
		return
	}
	p, ok := products[id]
	if !ok {
		c.String(http.StatusNotFound, "Product not found")
		return
	}
	c.IndentedJSON(http.StatusOK, p)

}

func get_port() string {
	port := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		port = ":" + val
	}
	return port
}
func main() {
	r := gin.Default()
	r.GET("/api/products", getProducts)
	r.GET("/api/products/:id", getProduct)
	r.Run(get_port())

}

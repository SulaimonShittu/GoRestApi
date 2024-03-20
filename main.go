package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type goods struct {
	ID           string  `json:"id"`
	ProductName  string  `json:"product-name"`
	Manufacturer string  `json:"manufacturer"`
	Price        float64 `json:"price"`
}

var goodslist = []goods{
	{ID: "1", ProductName: "Golden Penny Noodles", Manufacturer: "Erisco Foods", Price: 800},
	{ID: "2", ProductName: "Dangote Spaghetti", Manufacturer: "Dangote", Price: 750},
	{ID: "3", ProductName: "Honeywell seasoning goods", Manufacturer: "Honeywell", Price: 350},
}

func getGoods(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, goodslist)
}

func main() {
	router := gin.Default()
	router.GET("/goods", getGoods)
	router.POST("/goods", postGoods)
	router.GET("/albums/:id", getGoodByID)
	router.Run("localhost:8080")
}

func postGoods(c *gin.Context) {
	var newGood goods

	if err := c.BindJSON(&newGood); err != nil {
		return
	}

	goodslist = append(goodslist, newGood)
	c.IndentedJSON(http.StatusCreated, newGood)
}

func getGoodByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range goodslist {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "good not found"})
}

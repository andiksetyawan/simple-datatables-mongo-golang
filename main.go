package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/tables/serverside_table", func(c *gin.Context) {
		//c.Header("Access-Control-Allow-Origin", "*")

		var q Query
		if err := c.ShouldBindQuery(&q); err != nil {
			log.Println(err)
			return
		}

		var columns = []string{"A", "B", "C", "D"}
		collection := Mongo.Collection("emp")

		d := NewDatatable(collection, q, columns)
		d.generate()
		res := d.result()

		c.JSON(http.StatusOK, res)
	})

	r.Run(":9999")
}

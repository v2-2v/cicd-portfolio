package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.File("static/index.html")
	})

	r.Run(":8888") // ポート8080で起動
}

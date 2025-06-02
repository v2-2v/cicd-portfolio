package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/robots.txt", func(c *gin.Context) {
		c.Header("Content-Type", "text/plain")
		c.String(http.StatusOK, "User-agent: *\nDisallow: /")
	})

	r.GET("/", func(c *gin.Context) {
		c.File("static/index.html")
	})

	r.GET("/qr", func(c *gin.Context) {
		c.File("static/qr.png")
	})

	r.Run(":8888")
}

package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.SetHTMLTemplate(template.Must(template.ParseFiles(
		"templates/index.tmpl.html",
	)))
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"title": "TMPLチェック",
		})
	})

	r.Run(":8888")
}

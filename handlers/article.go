// handlers.article.go
package handlers

import (
	"net/http"
	"wiki-gin-gonic/models"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	articles := models.GetAllArticles()

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)

}

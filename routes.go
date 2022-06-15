package main

import (
	"wiki-gin-gonic/handlers"
)

func initializeRoutes() {

	// определение роута главной страницы
	router.GET("/", handlers.ShowIndexPage)

}

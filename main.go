package main

import (
	"users-books-api-testing/config"
	"users-books-api-testing/middlewares"
	"users-books-api-testing/routes"
)

func main() {
	config.InitDB()
	e := routes.New()

	// logger middleware
	middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}

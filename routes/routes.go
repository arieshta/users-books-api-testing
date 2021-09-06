package routes

import (
	"users-books-api-testing/config"
	"users-books-api-testing/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/login", controllers.LoginUserController)

	e.POST("/users", controllers.CreateUserController)

	// JWT Auth Group
	eJWT := e.Group("/jwt")
	eJWT.Use(middleware.JWT([]byte(config.SECRET_JWT)))
	eJWT.GET("/users", controllers.GetUsersController)
	eJWT.GET("/users/:id", controllers.GetUserByIdController)
	eJWT.PUT("/users/:id", controllers.UpdateUserByIdController)
	eJWT.DELETE("/users/:id", controllers.DeleteUserByIdController)

	e.POST("/books", controllers.AddBookController) //
	eJWT.GET("/books", controllers.GetBooksController)
	eJWT.GET("/books/:id", controllers.GetBookByIdController)
	eJWT.PUT("/books/:id", controllers.UpdateBookByIdController)    //
	eJWT.DELETE("/books/:id", controllers.DeleteBookByIdController) //

	return e
}

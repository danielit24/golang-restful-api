package main

import (
	"net/http"

	"github.com/danielit24/golang-restful-api/app"
	"github.com/danielit24/golang-restful-api/controller"
	"github.com/danielit24/golang-restful-api/exception"
	"github.com/danielit24/golang-restful-api/helper"
	"github.com/danielit24/golang-restful-api/middleware"
	"github.com/danielit24/golang-restful-api/repository"
	"github.com/danielit24/golang-restful-api/service"
	"github.com/go-playground/validator/v10"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	CategoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", CategoryController.FindAll)
	router.GET("/api/categories/:categoryId", CategoryController.FindById)
	router.POST("/api/categories", CategoryController.Create)
	router.PUT("/api/categories/:categoryId", CategoryController.Update)
	router.DELETE("/api/categories/:categoryId", CategoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

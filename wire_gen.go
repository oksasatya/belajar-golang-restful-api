// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"belajar-golang-restful-api/app"
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/middleware"
	"belajar-golang-restful-api/repository"
	"belajar-golang-restful-api/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"net/http"
)

import (
	_ "github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from injector.go:

func InitializedServer() *http.Server {
	categoryRepositoryImpl := repository.NewCategoryRepository()
	db := app.NewDB()
	validate := validator.New()
	categoryServiceImpl := service.NewCategoryService(categoryRepositoryImpl, db, validate)
	categoryControllerImpl := controller.NewCategoryController(categoryServiceImpl)
	router := app.NewRouter(categoryControllerImpl)
	authMiddleware := middleware.NewAuthMiddleware(router)
	server := NewServer(authMiddleware)
	return server
}

// injector.go:

var categorySet = wire.NewSet(repository.NewCategoryRepository, wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)), service.NewCategoryService, wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)), controller.NewCategoryController, wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)))

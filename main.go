package main

import (
	"github.com/bdxygy/go-rest-api/controller"
	"github.com/bdxygy/go-rest-api/repository"
	"github.com/bdxygy/go-rest-api/service"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	validate := validator.New()
	sql := NewDB()
	personRepositoryImpl := repository.NewPersonRepositoryImpl()
	personServiceImpl := service.NewPersonServiceImpl(sql, personRepositoryImpl, validate)
	personControllerImpl := controller.NewPersonControllerImpl(personServiceImpl)

	router.GET("/person", personControllerImpl.FindAll)
	router.POST("/person", personControllerImpl.Create)
	router.GET("/person/:personUUID", personControllerImpl.FindById)
	router.PUT("/person/:personUUID", personControllerImpl.Update)
	router.DELETE("/person/:personUUID", personControllerImpl.Delete)
}

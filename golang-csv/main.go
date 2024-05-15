package main

import (
	"golang-csv/config"
	"golang-csv/controller"
	"golang-csv/help"
	"golang-csv/modal"
	"golang-csv/repository"
	"golang-csv/router"
	"golang-csv/service"
	"net/http"
)

func main() {
	db := config.DatabaseConnection()
	db.AutoMigrate(&modal.FullDetails{})
	db.Table("employee_details").AutoMigrate(&modal.EmployeeDetails{})
	db.Table("manager_details").AutoMigrate(&modal.ManagerDetails{})
	userRepo := repository.NewUserRepoImp(db)
	Userservice := service.NewUserserviceImp(userRepo)
	UserController := controller.NewController(Userservice)
	router := router.Router(UserController)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	err := server.ListenAndServe()
	help.LisenError(err)
}

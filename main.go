package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"lms_try/database"
	"lms_try/handler"
	"lms_try/middleware"
	"lms_try/repository"
	"lms_try/router"
	"lms_try/service"
	"log"
	"net/http"
	"os"
)

func init() {
	// set log
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	// set godotenv
	if err := godotenv.Load(); err != nil {
		errMessage := fmt.Sprintf("error cant load env : %v", err.Error())
		logrus.Error(errMessage)
		log.Fatal(errMessage)
	}
}

func main() {
	// connect to MySQL database
	logrus.Info("this is from main program")

	db := database.DbConnection()

	// register repository
	userRepo := repository.NewUserRepository(db)

	// register Service
	userService := service.NewUserService(userRepo)

	// register handler
	userHandler := handler.NewuserHandler(userService)

	r := mux.NewRouter() //.PathPrefix("/api/v1/").Subrouter()
	r.Use(middleware.LoggerMiddleware)
	r.Use(middleware.AuthMiddleware)

	// generate user router
	router.GenerateUserRouter(r, userHandler)

	_ = http.ListenAndServe(":6000", r)

	logrus.WithFields(logrus.Fields{
		"package":  "main",
		"method":   "",
		"function": "main",
	}).Info("success running app")
}

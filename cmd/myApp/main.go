package main

import (
	"fmt"
	"net/http"

	"ninjaGo/internal/env"
	"ninjaGo/internal/handlers"
	"ninjaGo/internal/middleware"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("starting server")
	SetUp()

	router := mux.NewRouter()
	router.Use(mux.MiddlewareFunc(middleware.SessionMiddleware))

	router.HandleFunc("/login", handlers.LoginHandler).Methods(http.MethodPost)
	router.HandleFunc("/logout", handlers.Logout).Methods(http.MethodGet)
	router.HandleFunc("/health-check", handlers.HealthCheck).Methods(http.MethodGet)
	router.HandleFunc("/sign-up", handlers.SignUp).Methods(http.MethodPost)

	http.ListenAndServe(":8080", router)

}

func SetUp() {
	env.LoadEnvVariable()
}

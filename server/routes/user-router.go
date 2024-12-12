package routes

import (
	"stock-market-simulation/controller"

	"github.com/gorilla/mux"
)

func UserRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/user/signup", controller.Signup).Methods("POST")
	router.HandleFunc("/user/login", controller.Login).Methods("POST")

	return router
}

package routes

import (
	"stock-market-simulation/controller"

	"github.com/gorilla/mux"
)

func UserRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/signup", controller.Signup).Methods("POST")
	return router
}

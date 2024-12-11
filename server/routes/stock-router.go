package routes

import (
	"stock-market-simulation/controller"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

func StockRouter(rdb *redis.Client) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/stock/history", controller.GetStockHistoryHandler(rdb)).Methods("GET")

	return router
}

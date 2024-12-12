package routes

import (
	"stock-market-simulation/controller"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func StockRouter(rdb *redis.Client, db *mongo.Database) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/stock/history", controller.GetStockHistoryHandler(rdb)).Methods("GET")
	router.HandleFunc("/stock/buy", controller.BuyStockHandler(rdb, db)).Methods("POST")
	return router
}

package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func GetStockHistoryHandler(rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stock := r.URL.Query().Get("stock")
		if stock == "" {
			http.Error(w, "Missing stock parameter", http.StatusBadRequest)
			return
		}

		// Fetch historical prices for the stock from Redis sorted set
		data, err := rdb.ZRangeWithScores(context.Background(), stock+":history", 0, -1).Result()
		if err != nil {
			http.Error(w, "Failed to fetch stock history", http.StatusInternalServerError)
			return
		}

		history := []map[string]interface{}{}
		for _, entry := range data {
			price, err := strconv.ParseFloat(entry.Member.(string), 64) // Ensure it's parsed as float64
			if err != nil {
				http.Error(w, "Error parsing stock price", http.StatusInternalServerError)
				return
			}

			history = append(history, map[string]interface{}{
				"timestamp": entry.Score, // UNIX timestamp
				"price":     price,
			})
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(history) // Automatically writes 200 OK
	}
}

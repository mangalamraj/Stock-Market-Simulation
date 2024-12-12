package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"stock-market-simulation/models"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func BuyStockHandler(rdb *redis.Client, db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var buyRequest struct {
			Email    string `json:"email"`
			Stock    string `json:"stock"`
			Quantity int    `json:"quantity"`
		}

		if err := json.NewDecoder(r.Body).Decode(&buyRequest); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Get current stock price from Redis
		price, err := rdb.Get(context.Background(), buyRequest.Stock).Float64()
		if err != nil {
			http.Error(w, "Stock price not found", http.StatusNotFound)
			return
		}

		// Calculate total cost
		totalCost := price * float64(buyRequest.Quantity)

		// Get user from MongoDB
		collection := db.Collection("users")
		var user models.User
		filter := bson.M{"email": buyRequest.Email}

		if err := collection.FindOne(context.Background(), filter).Decode(&user); err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		// Check if user has enough balance
		if user.Balance < totalCost {
			http.Error(w, "Insufficient balance", http.StatusBadRequest)
			return
		}

		// Create a stock purchase entry
		stockPurchase := models.StockPurchase{
			Stock:    buyRequest.Stock,
			Quantity: buyRequest.Quantity,
			BuyPrice: price,
		}

		// Update user's balance and portfolio
		update := bson.M{
			"$inc": bson.M{"balance": -totalCost},
			"$set": bson.M{
				"portfolio": bson.A{stockPurchase},
			},
		}

		// If portfolio exists, use $push instead of $set
		if user.Portfolio != nil {
			update = bson.M{
				"$inc":  bson.M{"balance": -totalCost},
				"$push": bson.M{"portfolio": stockPurchase},
			}
		}

		log.Printf("Updating user %s with filter: %+v\n", buyRequest.Email, filter)
		log.Printf("Update operation: %+v\n", update)

		result, err := collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			log.Printf("MongoDB update error: %v\n", err)
			http.Error(w, "Failed to update user", http.StatusInternalServerError)
			return
		}

		log.Printf("Update result: MatchedCount: %d, ModifiedCount: %d\n",
			result.MatchedCount, result.ModifiedCount)

		// After successful MongoDB update, update the stock history
		timestamp := float64(time.Now().Unix())
		err = rdb.ZAdd(context.Background(), buyRequest.Stock+":history", redis.Z{
			Score:  timestamp,
			Member: strconv.FormatFloat(price, 'f', -1, 64),
		}).Err()
		if err != nil {
			log.Printf("Failed to update stock history: %v\n", err)
			// Don't return error since the purchase was successful
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Stock bought successfully",
			"cost":    totalCost,
			"price":   price,
		})
	}
}

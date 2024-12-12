package models

type StockPurchase struct {
	Stock    string  `bson:"stock"`
	Quantity int     `bson:"quantity"`
	BuyPrice float64 `bson:"buyPrice"`
}

type User struct {
	Email     string          `bson:"email" validate:"required"`
	Password  string          `bson:"password" validate:"required"`
	Name      string          `bson:"name" validate:"required"`
	Balance   float64         `bson:"balance"`
	Portfolio []StockPurchase `bson:"portfolio"`
}

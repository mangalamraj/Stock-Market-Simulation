package models

type User struct {
	Email     string             `bson:"email" validate:"required"`
	Password  string             `bson:"password" validate:"required"`
	Name      string             `bson:"name" validate:"required"`
	Balance   float64            `bson:"balance"`
	Portfolio []string           `bson:"portfolio"`
}

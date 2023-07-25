package graphql

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID     primitive.ObjectID `json:"id"`
	Amount float64            `json:"amount"`
	Date   time.Time          `json:"date"`
}

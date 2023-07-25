package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Transaction struct {
	ID     primitive.ObjectID `bson:"_id"`
	Amount float64            `bson:"amount"`
	Date   time.Time          `bson:"date"`
}

var collection *mongo.Collection

var transactionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Transaction",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
		},
		"amount": &graphql.Field{
			Type: graphql.Float,
		},
		"date": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"transaction": &graphql.Field{
			Type:        transactionType,
			Description: "Get single transaction",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := primitive.ObjectIDFromHex(params.Args["id"].(string))
				var transaction Transaction
				err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&transaction)
				if err != nil {
					log.Fatal(err)
				}
				return transaction, nil
			},
		},
		"list": &graphql.Field{
			Type:        graphql.NewList(transactionType),
			Description: "Get transaction list",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				cursor, err := collection.Find(context.Background(), bson.D{})
				if err != nil {
					log.Fatal(err)
				}
				var transactions []Transaction
				if err = cursor.All(context.Background(), &transactions); err != nil {
					log.Fatal(err)
				}
				return transactions, nil
			},
		},
	},
})

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("test").Collection("transactions")

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})

	http.Handle("/graphql", handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	}))

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}

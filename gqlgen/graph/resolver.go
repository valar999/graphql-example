package graph

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Resolver struct {
	transactionsCollection *mongo.Collection
}

func NewResolver() *Resolver {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	collection := client.Database("test").Collection("transactions")
	return &Resolver{transactionsCollection: collection}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Transactions(ctx context.Context) ([]*Transaction, error) {
	cursor, err := r.transactionsCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var transactions []*Transaction
	if err := cursor.All(ctx, &transactions); err != nil {
		return nil, err
	}
	return transactions, nil
}

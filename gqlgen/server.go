package main

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/valar999/graphql-example/gqlgen/graphql"
)

const defaultPort = "8080"

func main() {
	port := defaultPort

	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: graphql.NewResolver()}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("connect to http://localhost:" + port + " for GraphQL playground")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

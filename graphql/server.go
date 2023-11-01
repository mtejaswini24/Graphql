package main

import (
	"context"
	"fmt"
	"graphql/database"
	"graphql/graph"
	"graphql/service"
	"log"

	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	db, err := database.Open()
	if err != nil {
		fmt.Errorf("connecting to db %w", err)
	}
	pg, err := db.DB()
	if err != nil {
		fmt.Errorf("failed to get database connection %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = pg.PingContext(ctx) //verifies the database connection is there are not
	if err != nil {
		fmt.Errorf("db is not connected %w", err)
	}
	//initialize conn layer support
	ms, err := service.NewService(db)
	if err != nil {
		fmt.Errorf("initalizing connection failed")
	}
	srConn := service.NewStore(ms)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{S: srConn}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// func startApp() (*service.Conn, error) {
// 	//connection of DB
// 	db, err := database.Open()
// 	if err != nil {
// 		return &service.Conn{}, fmt.Errorf("connecting to db %w", err)
// 	}
// 	pg, err := db.DB()
// 	if err != nil {
// 		return &service.Conn{}, fmt.Errorf("failed to get database connection %w", err)
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
// 	defer cancel()

// 	err = pg.PingContext(ctx) //verifies the database connection is there are not
// 	if err != nil {
// 		return &service.Conn{}, fmt.Errorf("db is not connected %w", err)
// 	}
// 	//initialize conn layer support
// 	ms, err := service.NewService(db)
// 	if err != nil {
// 		return &service.Conn{}, err
// 	}
// 	return ms, nil
// }

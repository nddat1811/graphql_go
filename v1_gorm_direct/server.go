package main

import (
	"log"
	"net/http"
	"os"
	"v1_gorm_direct/config"
	"v1_gorm_direct/graph"
	"v1_gorm_direct/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/subosito/gotenv"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	err := gotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	conn := config.InitMysql()
	defer config.CloseConnectDB(conn)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

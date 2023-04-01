package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/site-tech/jaw-platform/ent"
	"github.com/site-tech/jaw-platform/graph"
)

const defaultPort = "8880"

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("error loading .env: ", err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	var entOptions []ent.Option
	entOptions = append(entOptions, ent.Debug())

	client, err := ent.Open("postgres", getDSN(os.Getenv("DATABASE_URL")), entOptions...)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	if os.Getenv("MIGRATE") == "true" {
		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed created schema resources: %v", err)
		}
	}

	srv := handler.NewDefaultServer(graph.NewSchema(client))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getDSN(url string) string {
	dsnFormat := "host=%v user=%v password=%v dbname=%v port=%v sslmode=disable"
	var host string
	var user string
	var word string
	var dbname string
	var port string

	s1 := strings.Split(url, "://")
	s2 := strings.Split(s1[1], ":")
	user = s2[0]
	s3 := strings.Split(s2[1], "@")
	word = s3[0]
	host = s3[1]
	s4 := strings.Split(s2[2], "/")
	port = s4[0]
	dbname = s4[1]

	return fmt.Sprintf(dsnFormat, host, user, word, dbname, port)
}

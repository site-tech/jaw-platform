package main

import (
	"context"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/site-tech/jaw-platform/ent"
	"github.com/site-tech/jaw-platform/graph"

	"github.com/newrelic/go-agent/v3/newrelic"

	kit_utils "github.com/sailsforce/gomicro-kit/utils"
)

const defaultPort = "8880"

func init() {
	if err := kit_utils.InitEnv(); err != nil {
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

	client, err := ent.Open("postgres", kit_utils.GetDSN(os.Getenv("DATABASE_URL")), entOptions...)
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

	// new relic
	newRelicApp, err := newrelic.NewApplication(
		newrelic.ConfigAppName(os.Getenv("NEW_RELIC_APP_NAME")),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_APP_LICENSE")),
	)
	if err != nil {
		log.Fatalf("new relic init failed: %v", err)
	}

	srv := handler.NewDefaultServer(graph.NewSchema(client))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle(newrelic.WrapHandle(newRelicApp, "/query", srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"ariga.io/entcache"
	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/site-tech/jaw-platform/ent"
	"github.com/site-tech/jaw-platform/graph"
)

const defaultPort = "8880"

var entCacheDriver *entcache.Driver
var ttl time.Duration = time.Minute * 1

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

	// CORS and router
	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"GET,POST,PUT,DELETE,OPTIONS"},
		AllowCredentials:   true,
		Debug:              true,
		OptionsPassthrough: true,
	}).Handler)

	// set TTL
	t, err := strconv.ParseInt(os.Getenv("CACHE_TTL"), 10, 0)
	if err != nil {
		log.Println("cache_ttl env var err: ", err)
	} else {
		ttl = time.Duration(t) * time.Minute
	}

	// create cached ent client
	ctx := context.Background()
	client := createCachedEntClient(ctx)
	defer client.Close()

	// Run the auto migration tool.
	if os.Getenv("MIGRATE") == "true" {
		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed created schema resources: %v", err)
		}
	}

	srv := handler.NewDefaultServer(graph.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})
	//gql-srv-response-middleware: context-level caching
	srv.AroundResponses(func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
		if op := graphql.GetOperationContext(ctx).Operation; op != nil && op.Operation == ast.Query {
			ctx = entcache.NewContext(ctx)
		}
		return next(ctx)
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	// endpoint for chache stats
	router.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		if entCacheDriver != nil {
			stat := entCacheDriver.Stats()
			fmt.Fprintf(w, "cache stats (gets: %d, hits: %d, errors: %d)\n", stat.Gets, stat.Hits, stat.Errors)
		} else {
			fmt.Fprintln(w, "cache mode is not enabled")
		}
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
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

func createCachedEntClient(ctx context.Context) *ent.Client {
	return ent.NewClient([]ent.Option{
		ent.Debug(),
		ent.Driver(newDialectDriver(ctx)),
	}...)
}

func newDialectDriver(ctx context.Context) dialect.Driver {
	dbDSN := fmt.Sprintf("%s %s%s", getDSN(os.Getenv("DATABASE_URL")), "options=--search_path=", os.Getenv("DATABASE_SCHEMA"))
	db, err := sql.Open("postgres", dbDSN)
	if err != nil {
		log.Fatal("opening database: ", err)
	}

	// launch redis client
	fullURL := os.Getenv("REDIS_URL")
	pwd := strings.Split(fullURL, "@")[0][9:]
	addr := strings.Split(fullURL, "@")[1]
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
	})
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatal("redis ping: ", err)
	}

	entCacheDriver = entcache.NewDriver(
		dialect.Debug(db),
		entcache.TTL(ttl),
		entcache.Levels(
			entcache.NewLRU(256),
			entcache.NewRedis(redisClient),
		),
	)

	return entCacheDriver
}

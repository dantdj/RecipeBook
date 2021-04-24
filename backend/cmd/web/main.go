package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/dantdj/RecipeBook/internal/config"
	"github.com/dantdj/RecipeBook/pkg/models/mongooperations"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	recipes  *mongooperations.RecipeDb
}

func main() {
	config.LoadConfig()

	addr := flag.String("addr", "localhost:4000", "Address to run server on")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client := getDatabaseClient(ctx)

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		recipes:  &mongooperations.RecipeDb{Client: client},
	}

	server := http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		WriteTimeout: 1 * time.Second,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := server.ListenAndServe()
	errorLog.Fatal(err)
}

func getDatabaseClient(ctx context.Context) (client *mongo.Client) {
	clientOptions := options.Client().ApplyURI(config.Configuration.Mongo.ConnectionString).SetDirect(true)

	client, err := mongo.NewClient(clientOptions)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("unable to initialize connection %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("unable to connect %v", err)
	}

	return client
}

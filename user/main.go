package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"adomeit.xyz/user/ent"
	"entgo.io/ent/dialect"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open the database connection
	client := SetupClient()
	defer client.Close()

	// Run auto migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "Failed creating schema: %v\n", err)
		os.Exit(1)
	}

	r := gin.Default()
	manager := NewAuthManager(getenv("JWT_SECRET", "NON_SECRET_DEFAULT"))
	shardMap, err := GetShardMap()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid shard map: %v\n", err)
		os.Exit(1)
	}

	NewUserController(r, client, manager, &shardMap)

	// Set up the routes available in the API
	r.Run()
}

const clientTimeout = 2 * 60 * time.Second

func SetupClient() *ent.Client {

	db_url := os.Getenv("DATABASE_URL")
	var client *ent.Client
	var e error
	if len(db_url) == 0 {
		client, e = ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
		if e != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", e)
			os.Exit(1)
		}
		return client
	}
	sleepTimer := time.Second * 3
	for {
		db, err := sql.Open("pgx", db_url)
		if err == nil && db.Ping() == nil {

			// Wrap the database connection in the ent driver and create the client
			drv := entsql.OpenDB(dialect.Postgres, db)

			return ent.NewClient(ent.Driver(drv))
		}
		sleepTimer = sleepTimer * 2
		if sleepTimer > clientTimeout {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v (Retries timed out) \n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v (Retry in %v seconds) \n", err, sleepTimer)
		time.Sleep(sleepTimer)
	}
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

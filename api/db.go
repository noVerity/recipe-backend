package main

import (
	"adomeit.xyz/recipe/ent"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"os"
	"time"
)

const clientTimeout = 2 * 60 * time.Second

func SetupClient() *ent.Client {
	dbUrl := os.Getenv("DATABASE_URL")
	var client *ent.Client
	var e error
	if len(dbUrl) == 0 {
		client, e = ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
		if e != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", e)
			os.Exit(1)
		}
		return client
	}
	sleepTimer := time.Second * 3
	for {
		db, err := sql.Open("pgx", dbUrl)
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

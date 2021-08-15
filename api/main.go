package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"adomeit.xyz/recipe/ent"
	"entgo.io/ent/dialect"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	db_url = "postgres://postgres:mysecretpassword@192.168.1.118:5432/recipe?sslmode=disable"
)

func main() {
	// Open the database connection
	db, err := sql.Open("pgx", getenv("DATABASE_URL", db_url))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// Wrap the database connection in the ent driver and create the client
	drv := entsql.OpenDB(dialect.Postgres, db)

	client := ent.NewClient(ent.Driver(drv))
	defer client.Close()

	// Run auto migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "Failed creating schema: %v\n", err)
		os.Exit(1)
	}

	manager := NewAuthManager(getenv("JWT_SECRET", "NON_SECRET_DEFAULT"))

	// Set up the routes available in the API
	r := SetupRouter(client, gin.Default(), manager)
	r.Run()
}


func SetupRouter(client *ent.Client, r *gin.Engine, auth *AuthManager) *gin.Engine {
	NewUserController(r, client, auth)
	NewIngredientController(r, client, auth)
	NewRecipeController(r, client, auth)

	return r
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

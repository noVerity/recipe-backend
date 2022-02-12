package main

import (
	"adomeit.xyz/recipe/core"
	"adomeit.xyz/recipe/frontend/rest"
	"adomeit.xyz/recipe/mq"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/mattn/go-sqlite3"
	"os"
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

	shard := getenv("SHARD", "0")

	queue := mq.NewMQ(
		getenv("CLOUDAMQP_URL", "amqp://guest:guest@localhost:5672/"),
		shard,
		getenv("APP_OUT_QUEUE", "ingredients_lookup"),
		getenv("APP_IN_QUEUE", "ingredients_results"),
	)
	defer queue.Close()

	recipeCore := core.NewRecipeCore(client, queue)
	ingredientCore := core.NewIngredientCore(client, queue)

	err := ingredientCore.AcceptIngredientResults(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed connecting to MQ: %v\n", err)
		os.Exit(1)
	}

	manager := rest.NewAuthManager(getenv("JWT_SECRET", "NON_SECRET_DEFAULT"))
	// Set up the routes available in the API
	r := rest.SetupRouter(gin.Default(), manager, recipeCore, ingredientCore)
	r.Run()
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

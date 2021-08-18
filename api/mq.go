package main

import (
	"context"
	"encoding/json"
	"log"

	"adomeit.xyz/recipe/ent"
	"github.com/streadway/amqp"
)

type IngredientMQRequest struct {
	RecipeID   int    `json:"recipeId"`
	SearchTerm string `json:"searchTerm"`
}

type IngredientMQResult struct {
	IngredientMQRequest

	Name          string  `json:"name"`
	Calories      float32 `json:"calories"`
	Protein       float32 `json:"protein"`
	Fat           float32 `json:"fat"`
	Carbohydrates float32 `json:"carbohydrates"`
}

func RequestIngredients(entries []IngredientEntry, recipeID int) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		getenv("APP_OUT_QUEUE", "ingredients_lookup"), // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	for _, entry := range entries {
		request := IngredientMQRequest{
			RecipeID:   recipeID,
			SearchTerm: entry.Name,
		}

		message, err := json.Marshal(request)
		failOnError(err, "Failed to marshal a message")

		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        []byte(message),
			})
		failOnError(err, "Failed to publish a message")
	}
}

func AcceptIngredientResults(client *ent.Client) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		getenv("APP_IN_QUEUE", "ingredients_results"), // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var message IngredientMQResult
			json.Unmarshal(d.Body, &message)
			log.Printf("Received a message: %v", message)
			CreateIngredientFromMessage(client, message)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func CreateIngredientFromMessage(client *ent.Client, message IngredientMQResult) {
	ctx := context.Background()
	ingredient, err := client.Ingredient.Create().
		SetName(message.SearchTerm).
		SetCalories(message.Calories).
		SetFat(message.Fat).
		SetCarbohydrates(message.Carbohydrates).
		SetProtein(message.Protein).
		SetSource(message.Name).
		Save(ctx)

	if err != nil {
		return
	}

	client.Recipe.UpdateOneID(message.RecipeID).
		AddIngredients(ingredient).
		Save(ctx)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

package main

import (
	"context"
	"encoding/json"
	"log"
	"math"
	"sync"
	"time"

	"adomeit.xyz/recipe/ent"
	"github.com/streadway/amqp"
)

type MQ struct {
	mu         sync.Mutex
	connection *amqp.Connection
	channel    *amqp.Channel
	location   string
	delay      time.Duration

	LookupQueue *amqp.Queue
	ResultQueue *amqp.Queue
}

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

func NewMQ(connectionString string) *MQ {
	mq := &MQ{sync.Mutex{}, nil, nil, connectionString, 0, nil, nil}

	return mq
}

func (mq *MQ) Connect() error {
	mq.mu.Lock()
	if mq.connection == nil || mq.connection.IsClosed() {
		conn, err := amqp.Dial(mq.location)

		if err != nil {
			// Try to connect again, but don't DDOS our own service
			go mq.reconnect()
			// Call lock here again to wait for the reconnect to unlock eventually
			mq.mu.Lock()
		} else {
			log.Print("Connected to message queue")
			mq.connection = conn
		}
	}
	defer mq.mu.Unlock()

	if mq.channel == nil {
		ch, err := mq.connection.Channel()

		if err != nil {
			return err
		}

		mq.channel = ch

		mq.channelDeclare()
	}

	return nil
}

func (mq *MQ) reconnect() {
	// Try to connect again, but don't DDOS our own service
	// We are starting with 10 seconds and the gradually back off to a 30 minute interval
	if mq.delay == 0 {
		mq.delay = time.Second * 10
	} else {
		mq.delay = time.Duration(int64(math.Min(float64(mq.delay*2), float64(time.Minute*30))))
	}
	log.Printf("Failed to connect to message queue trying again in %v", mq.delay)
	time.Sleep(mq.delay)
	conn, err := amqp.Dial(mq.location)
	if err != nil {
		go mq.reconnect()
		return
	}
	mq.connection = conn
	mq.delay = 0
	log.Print("Connected to message queue")
	mq.mu.Unlock()
}

func (mq *MQ) Close() {
	mq.mu.Lock()
	defer mq.mu.Unlock()
	if mq.channel != nil {
		mq.channel.Close()
	}

	if mq.connection != nil && !mq.connection.IsClosed() {
		mq.connection.Close()
	}
}

func (mq *MQ) channelDeclare() error {

	mq.mu.Lock()
	defer mq.mu.Unlock()
	lookupQueue, err := mq.channel.QueueDeclare(
		getenv("APP_OUT_QUEUE", "ingredients_lookup"), // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		return err
	}

	mq.LookupQueue = &lookupQueue

	resultQueue, err := mq.channel.QueueDeclare(
		getenv("APP_OUT_QUEUE", "ingredients_lookup"), // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		return err
	}

	mq.ResultQueue = &resultQueue

	return nil
}

func (mq *MQ) RequestIngredients(entries []IngredientEntry, recipeID int) {
	mq.Connect()

	for _, entry := range entries {
		request := IngredientMQRequest{
			RecipeID:   recipeID,
			SearchTerm: entry.Name,
		}

		message, err := json.Marshal(request)
		failOnError(err, "Failed to marshal a message")

		err = mq.channel.Publish(
			"",                  // exchange
			mq.LookupQueue.Name, // routing key
			false,               // mandatory
			false,               // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        []byte(message),
			})
		failOnError(err, "Failed to publish a message")
	}
}

func (mq *MQ) AcceptIngredientResults(client *ent.Client) {

	mq.Connect()

	msgs, err := mq.channel.Consume(
		mq.ResultQueue.Name, // queue
		"",                  // consumer
		true,                // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,                 // args
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

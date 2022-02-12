package mq

import (
	"encoding/json"
	"github.com/noVerity/gofromto"
	"log"
	"math"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

type MQ struct {
	mu          sync.Mutex
	connection  *amqp.Connection
	channel     *amqp.Channel
	location    string
	delay       time.Duration
	shard       string
	lookupQueue string
	resultQueue string

	LookupQueue *amqp.Queue
	ResultQueue *amqp.Queue
}

type IngredientMQRequest struct {
	RecipeID   string `json:"recipeId"`
	Shard      string `json:"shard"`
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

func NewMQ(connectionString string, shard string, lookupQueue string, resultQueue string) *MQ {
	mq := &MQ{sync.Mutex{}, nil, nil, connectionString, 0, shard, lookupQueue, resultQueue, nil, nil}

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
		mq.lookupQueue, // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)

	if err != nil {
		return err
	}

	mq.LookupQueue = &lookupQueue

	resultQueue, err := mq.channel.QueueDeclare(
		mq.resultQueue+mq.shard, // name
		false,                   // durable
		false,                   // delete when unused
		false,                   // exclusive
		false,                   // no-wait
		nil,                     // arguments
	)

	if err != nil {
		return err
	}

	mq.ResultQueue = &resultQueue

	return nil
}

func (mq *MQ) RequestIngredients(entries []gofromto.Measure, recipeID string) {
	mq.Connect()

	for _, entry := range entries {
		request := IngredientMQRequest{
			RecipeID:   recipeID,
			SearchTerm: entry.Name,
			Shard:      mq.shard,
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

func (mq *MQ) AcceptIngredientResults() (chan IngredientMQResult, error) {
	ingredientResults := make(chan IngredientMQResult)

	err := mq.Connect()

	if err != nil {
		return ingredientResults, err
	}

	msgs, err := mq.channel.Consume(
		mq.ResultQueue.Name, // queue
		"",                  // consumer
		true,                // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,                 // args
	)

	if err != nil {
		close(ingredientResults)
		return ingredientResults, err
	}

	go func() {
		for d := range msgs {
			var message IngredientMQResult
			err := json.Unmarshal(d.Body, &message)
			if err != nil {
				log.Printf("Invalid message")
				continue
			}
			log.Printf("Received a message: %v", message)
			ingredientResults <- message
		}
	}()

	return ingredientResults, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

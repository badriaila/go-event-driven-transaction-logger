package main

import (
	"context"
	"encoding/json"
	"log"

	"transaction-logger/database"
	"transaction-logger/models"

	"github.com/segmentio/kafka-go"
)

func main() {
	database.ConnectDB()
	reader := kafka.NewReader(kafka.ReaderConfig{ // Kafka reader configuration
		// Adjust the broker address, topic, and group ID as needed
		Brokers: []string{"localhost:9092"},
		Topic:   "transactions",
		GroupID: "transaction-logger-consumer-group-1",
	})

	defer reader.Close()

	log.Println("Starting Transaction Consumer...")

	for {
		m, err := reader.ReadMessage(context.Background()) // Read messages from the Kafka topic

		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		var txn models.Transaction
		err = json.Unmarshal(m.Value, &txn) // Unmarshal the message value into a Transaction struct

		if err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			continue
		}
		log.Printf("Consumed transaction: %s", txn.ID)
		res := database.DB.Create(&txn)

		if res.Error != nil {
			log.Printf("Error saving transaction to database: %v", res.Error)
			continue
		}
		log.Printf("Transaction saved to database: %s", txn.ID)
	}

}

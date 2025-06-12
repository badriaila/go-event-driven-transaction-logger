package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/segmentio/kafka-go"
)

type Transaction struct {
	ID        string  `json:"id"`
	Type      string  `json:"type"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
	Status    string  `json:"status"`
	Timestamp string  `json:"timestamp"`
}

func randomTransaction() Transaction {
	types := []string{"CREDIT", "DEBIT"}
	statuses := []string{"PENDING", "COMPLETED", "FAILED"}
	currencies := []string{"USD", "EUR", "INR"}

	return Transaction{
		ID:        "txn_" + randomString(8),
		Type:      types[rand.Intn(len(types))],
		Amount:    float64(rand.Intn(1000)) / 100.0,
		Currency:  currencies[rand.Intn(len(currencies))],
		Status:    statuses[rand.Intn(len(statuses))],
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)

	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func main() {
	writer := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "transactions",
		Balancer: &kafka.LeastBytes{},
	}

	defer writer.Close()

	log.Println("Starting Transaction Producer...")

	for i := 0; i < 10; i++ {
		txn := randomTransaction()
		value, _ := json.Marshal(txn)

		err := writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(txn.ID),
				Value: value,
			},
		)

		if err != nil {
			log.Printf("Failed to write message: %v\n", err)
			continue
		} else {
			log.Printf("Produced transaction: %s\n", txn.ID)
		}
		time.Sleep(1 * time.Second) // To simulate some delay between transactions

	}

	log.Println("Transaction Producer finished.")
	log.Println("Exiting...")
}

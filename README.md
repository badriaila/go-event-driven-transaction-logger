# 🧾 Event-Driven Transaction Logger (Fintech)

A Go-based microservice that simulates a real-world **Fintech Transaction Logging System**. It uses **Kafka** to receive transaction events, persists them in a **PostgreSQL** database, and exposes a **REST API** to query the transaction logs.

---

## 🚀 What I am Building

---

## 🎯 Features

- ✅ Kafka Producer to simulate mock transactions
- ✅ Kafka Consumer to process and validate incoming events
- ✅ PostgreSQL database for persistent logging
- ✅ REST API to:
  - List all transactions
  - Get transaction by ID
  - Filter by status or date
- ✅ Docker-based setup for local development

---

## 🧰 Tech Stack

| Layer           | Technology                     |
|----------------|---------------------------------|
| Programming     | Golang                         |
| Messaging Queue | Kafka + Zookeeper              |
| Database        | PostgreSQL                     |
| API Framework   | Gin (or Chi, TBA)              |
| Kafka Client    | segmentio/kafka-go             |
| ORM/DB          | GORM (or sqlx / native `sql`)  |
| DevOps          | Docker + Docker Compose        |

---

---

## 🏃 Getting Started

```bash
# Step 1: Start dependencies
docker-compose up -d

# Step 2: Run the producer (send sample transactions)
go run producer/main.go

# Step 3: Run the consumer (listen + write to DB)
go run consumer/main.go

# Step 4: Query transactions via API
curl http://localhost:8080/transactions


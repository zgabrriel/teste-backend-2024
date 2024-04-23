package products

import (
	"context"
	"encoding/json"
	"log"
	"ms-go/app/models"
	"time"

	"github.com/segmentio/kafka-go"
)

func Upsert(data models.Product) {
	productJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err.Error())
	}

	conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:29092", "go-to-rails", 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(productJSON)},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

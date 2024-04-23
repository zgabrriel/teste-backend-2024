package consumers

import (
	"context"
	"encoding/json"
	"fmt"
	"ms-go/db"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewKafkaReader() *kafka.Reader {
	connection := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:29092"},
		Topic:   "rails-to-go",
	})
	return connection
}

func RailsConsumer() {
	connection := NewKafkaReader()
	defer connection.Close()

	for {
		message, err := connection.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("failed to read message:", err)
		}

		var data map[string]interface{}
		if err := json.Unmarshal(message.Value, &data); err != nil {
			fmt.Println("failed decoding message JSON:", err)
			return
		}

		db.Connection().UpdateOne(context.TODO(), bson.M{"id": data["id"]}, bson.M{"$set": data}, options.Update().SetUpsert(true))

		defer db.Disconnect()
	}

}

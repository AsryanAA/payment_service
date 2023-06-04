package utils

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func WriteApplicationInQueue(app []byte) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "payment_service_topic", 0)

	if err != nil {
		log.Fatal("Error while connect kafka service", err)
	}

	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))

	conn.WriteMessages(kafka.Message{Value: app})

	if err := conn.Close(); err != nil {
		log.Fatal("Error while close producer", err)
	}
}

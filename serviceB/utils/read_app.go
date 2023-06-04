package utils

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func ReadApplicationOutQueue() {
	/*conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "payment_service_topic", 0)

	if err != nil {
		log.Fatal("Error while connect kafka service", err)
	}

	conn.SetReadDeadline(time.Now().Add(time.Second * 10))

	batch := conn.ReadBatch(1e3, 1e9) // 1 MB to 1 GB
	bytes := make([]byte, 191)

	for {
		n, err := batch.Read(bytes)
		if err != nil {
			break
		}
		fmt.Println(string(bytes[:n]))
	}*/
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "payment_service_topic",
		Partition: 0,
		GroupID:   "consumer-id",
		MaxBytes:  10e6, // 10MB
	})
	r.SetOffset(0)

	/*for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}*/
	ctx := context.Background()
	for {
		message, err := r.FetchMessage(ctx)
		if err != nil {
			break
		}
		//fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		if err := r.CommitMessages(ctx, message); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
		CheckApplication(message.Value)
	}
}

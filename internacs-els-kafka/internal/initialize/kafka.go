package initialize

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"internacs-els-kafka/global"
	"internacs-els-kafka/models"
	"internacs-els-kafka/utils"
	"log"
	"os"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	kafkaURL   = os.Getenv("KAFKA_URL")
	kafkaTopic = os.Getenv("KAFKA_TOPIC")
)

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		StartOffset:    kafka.FirstOffset,
	})
}

func registerDebeziumConsumer(groupId string) {
	kafkaGroupId := fmt.Sprintf("consumer-group-%s", groupId)
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)
	defer reader.Close()

	fmt.Printf("Consumer %s start\n", groupId)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Consumer %s error: %s\n", groupId, err.Error())
			continue
		}
		var event models.DebeziumEvent
		json.Unmarshal([]byte(m.Value), &event)

		if event.Op == "c" {
			hotel, err := utils.ConvertAfterToHotel(event)
			if err != nil {
				fmt.Printf("Error converting after to hotel: %v\n", err)
				return
			}

			data, err := json.Marshal(hotel)
			if err != nil {
				fmt.Printf("Error marshalling hotel: %v\n", err)
				return
			}

			global.Els.Index(index, bytes.NewReader(data))
		}

	}
}

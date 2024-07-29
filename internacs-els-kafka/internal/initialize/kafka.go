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
	"sync"
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
		MaxBytes:       50e6, // 50MB
		CommitInterval: time.Second,
		StartOffset:    kafka.FirstOffset,
	})
}

func registerDebeziumConsumer(groupId string) {
	kafkaGroupId := fmt.Sprintf("consumer-group-%s", groupId)
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)
	defer reader.Close()

	fmt.Printf("Consumer %s start\n", groupId)

	var wg sync.WaitGroup
	messageChannel := make(chan kafka.Message, 1000)

	// Increase number of goroutines based on your system's capability
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for m := range messageChannel {
				var event models.DebeziumEvent
				err := json.Unmarshal(m.Value, &event)
				if err != nil {
					log.Printf("Error unmarshalling message: %v", err)
					continue
				}

				if event.Op == "c" {
					hotel, err := utils.ConvertAfterToHotel(event)
					if err != nil {
						fmt.Printf("Error converting after to hotel: %v\n", err)
						continue
					}

					data, err := json.Marshal(hotel)
					if err != nil {
						fmt.Printf("Error marshalling hotel: %v\n", err)
						continue
					}

					// Use bulk indexing for better performance
					_, err = global.Els.Index(index, bytes.NewReader(data))
					if err != nil {
						fmt.Printf("Error indexing document: %v\n", err)
					}
				}
			}
		}()
	}

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Consumer %s error: %s\n", groupId, err.Error())
			continue
		}

		messageChannel <- m
	}

	close(messageChannel)
	wg.Wait()
}

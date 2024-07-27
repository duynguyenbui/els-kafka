package initialize

import (
	"fmt"
	"internacs-els-kafka/global"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	index = os.Getenv("ELASTIC_INDEX")
)

func InitElasticSearch() {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"http://elasticsearch:9200",
		},
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Uncomment if you want to check connection
	res, err := client.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error response from Elasticsearch: %s", res.Status())
	}

	log.Printf("Status: %s", res.Status())

	fmt.Println("Elasticsearch connection established successfully")

	global.Els = client
}

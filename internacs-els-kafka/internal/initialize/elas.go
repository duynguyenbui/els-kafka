package initialize

import (
	"fmt"
	"internacs-els-kafka/global"
	"log"
	"os"

	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
)

var (
	index = os.Getenv("ELASTIC_INDEX")
)

// TODO: Change it to be configurable
func InitElasticSearch() {
	cfg := elasticsearch7.Config{
		Addresses: []string{
			"http://elasticsearch:9200",
		},
		Username: "elastic",
		Password: "elasticpw",
	}
	es, err := elasticsearch7.NewClient(cfg)

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Uncomment if you want to check connection
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error response from Elasticsearch: %s", res.Status())
	}

	log.Printf("Status: %s", res.Status())

	fmt.Println("Elasticsearch connection established successfully")

	global.Els = es
}

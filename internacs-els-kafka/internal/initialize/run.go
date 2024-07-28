package initialize

func Run() {
	// Init elastic search
	InitElasticSearch()

	// Init Redis
	InitRedis()

	// Init Postgres
	InitPostgres()

	// Init Kafka
	go registerDebeziumConsumer("elsgo")

	// go data.Seed()
}

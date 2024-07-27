curl -i -X PUT -H  "Content-Type:application/json" \
    http://localhost:8083/connectors/debezium-kafkaconnect/config \
    -d @docker/kafka/connects/debezium-kafkaconnect.json
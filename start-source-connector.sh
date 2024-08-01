curl -i -X PUT -H  "Content-Type:application/json" \
    http://localhost:8083/connectors/debezium-kafkaconnect-hotels/config \
    -d @docker/kafka/connects/debezium-kafkaconnect-hotels.json

curl -i -X PUT -H  "Content-Type:application/json" \
    http://localhost:8083/connectors/debezium-kafkaconnect-regions/config \
    -d @docker/kafka/connects/debezium-kafkaconnect-regions.json
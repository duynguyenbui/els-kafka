package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/klauspost/compress/zstd"
)

// Define the Hotel struct based on your JSON data structure
type Hotel struct {
	ID      string `json:"id"`
	Address string `json:"address"`
	// Add other fields based on your JSON structure
}

func ParseData(fileName string) {
	filePath := fileName

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	zstdReader, err := zstd.NewReader(file)
	if err != nil {
		log.Fatalf("Failed to create zstd reader: %v", err)
	}
	defer zstdReader.Close()

	reader := bufio.NewReader(zstdReader)
	i := 0

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		line = line[:len(line)-1]

		var record Hotel
		if err := json.Unmarshal([]byte(line), &record); err != nil {
			log.Printf("Failed to unmarshal line: %v", err)
			continue
		}

		// do stuff with the record
		log.Printf("record %d is %s and the address is %s", i, record.ID, record.Address)
		i++
	}
}

func main() {
	fileName := "partner_feed_en_v3_minimal.jsonl.zst"
	ParseData(fileName)
}

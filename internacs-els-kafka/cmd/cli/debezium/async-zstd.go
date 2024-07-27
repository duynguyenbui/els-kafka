/*
ETG API provides hotel's static data dump in .zstd format.
You can find more about the dump structure and the format in our documentation - https://docs.emergingtravel.com/#0b55c99a-7ef0-4a18-bbfe-fd1bdf35d08e
Please note that uncompressed data could be more than 20GB.
Below is an example of how to handle such large archive.
For decompression, we will use the zstd package which you can install using the command
> go get github.com/DataDog/zstd
The function takes the path to the archive file,
splits the whole file by 16MB chunks,
extracts objects line by line (each line contains one hotel in JSON format),
and unmarshals them into Golang structs which you can use in your inner logic.
The main difference between async and sync modes is the time of processing:
async is faster as each chunk will be handled asynchronously.
*/
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"math"
	"os"

	"golang.org/x/sync/semaphore"

	//go get github.com/DataDog/zstd
	"github.com/DataDog/zstd"
)

// Raw is a storage with raw data
type Raw struct {
	firstLine []byte
	lastLine  []byte
}

// copySlice helps copy raw data without memory leak
func copySlice(slice []byte) []byte {
	copiedSlice := make([]byte, len(slice))
	for i, v := range slice {
		copiedSlice[i] = v
	}
	return copiedSlice
}

// processHotel works stuff with raw hotel byte data
func processHotel(hotelRaw []byte) {
	var hotel Hotel
	err := json.Unmarshal(hotelRaw, &hotel)
	if err != nil {
		log.Println(err)
	}
	// do stuff with the hotel
	log.Printf("current hotel is %s", hotel.Address)
}

// processChunk works with raw batches
func processChunk(chunk []byte, sem *semaphore.Weighted, rawChan chan Raw) {
	defer sem.Release(1)
	lines := bytes.Split(chunk, []byte("\n"))

	// Ensure there are at least two lines (first and last) to avoid out-of-bounds access
	if len(lines) > 1 {
		rawChan <- Raw{
			firstLine: copySlice(lines[0]),
			lastLine:  copySlice(lines[len(lines)-1]),
		}
		for _, line := range lines[1 : len(lines)-1] {
			processHotel(line)
		}
	} else if len(lines) == 1 {
		// Handle the case where there's only one line in the chunk
		rawChan <- Raw{
			firstLine: copySlice(lines[0]),
			lastLine:  copySlice(lines[0]),
		}
		processHotel(lines[0])
	} else {
		// No lines found, send an empty Raw
		rawChan <- Raw{}
	}
}

// processRawHotels processes raw hotel data
func processRawHotels(raws []Raw) {
	for i, r := range raws {
		if i == 0 {
			processHotel(r.firstLine)
			continue
		}
		data := append(raws[i-1].lastLine, r.firstLine...)
		processHotel(data)
	}
}
func parseDump(filename string) {
	// open zst file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// make zstd reader
	reader := zstd.NewReader(file)
	// we will work the file by 16mb chunk
	bufferSize := make([]byte, int(math.Pow(2, 24)))
	// with weighted semaphore by max 10 async goroutines
	ctx := context.Background()
	var sem = semaphore.NewWeighted(int64(10))
	// and make the storage and the transport for raw data
	// the firstLine and the lastLine lines from a chunk
	rawData := make([]Raw, 0)
	rawChan := make(chan Raw)
	isFinished := false
	for {
		if isFinished {
			break
		}
		n, readErr := reader.Read(bufferSize)
		if readErr != nil && readErr != io.EOF {
			log.Fatal(readErr)
		}
		// stop loop if EOF
		if readErr == io.EOF {
			isFinished = true
		}
		// slices are pointers
		// copy it
		rawReadData := bufferSize[:n]
		actualLine := make([]byte, len(rawReadData))
		copy(actualLine, rawReadData)
		// all JSON files split by the new line char "\n"
		// try to read one by one
		_ = sem.Acquire(ctx, 1)
		go processChunk(actualLine, sem, rawChan)
		rawData = append(rawData, <-rawChan)
	}
	processRawHotels(rawData)
}

// func main() {
// 	parseDump("partner_feed_en_v3_minimal.jsonl.zst")
// }

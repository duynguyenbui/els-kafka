package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"time"

	"github.com/duynguyenbui/async-reader/models"
	_ "github.com/lib/pq"

	"github.com/DataDog/zstd"
	"golang.org/x/sync/semaphore"
)

// Raw is a storage with raw data
type Raw struct {
	firstLine []byte
	lastLine  []byte
}

var db *sql.DB

// copySlice helps copy raw data without memory leak
func copySlice(slice []byte) []byte {
	copiedSlice := make([]byte, len(slice))
	copy(copiedSlice, slice)
	return copiedSlice
}

// processHotel works stuff with raw hotel byte data
func processHotel(hotelRaw []byte) {
	var hotel models.Hotel
	err := json.Unmarshal(hotelRaw, &hotel)

	if err == nil {
		time.Sleep(time.Millisecond * 1)
		insertHotel(db, hotel)
	}

}

func processChunk(chunk []byte, sem *semaphore.Weighted, rawChan chan Raw) {
	defer sem.Release(1)
	lines := bytes.Split(chunk, []byte("\n"))
	if len(lines) == 0 {
		log.Println("Empty chunk received, skipping")
		return
	}
	rawChan <- Raw{
		firstLine: copySlice(lines[0]),
		lastLine:  copySlice(lines[len(lines)-1]),
	}
	for i := 1; i < len(lines)-1; i++ {
		if len(lines[i]) == 0 {
			log.Printf("Skipping empty line at index %d", i)
			continue
		}
		processHotel(lines[i])
	}
}

func processRawHotels(raws []Raw) {
	for i := range raws {
		if i == 0 {
			processHotel(raws[i].firstLine)
		} else {
			data := append(raws[i-1].lastLine, raws[i].firstLine...)
			processHotel(data)
		}
		// Process the complete last line if it's not the first chunk
		if len(raws[i].lastLine) > 0 && i < len(raws)-1 {
			processHotel(raws[i].lastLine)
		}
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
	defer reader.Close()

	// we will work the file by 16mb chunk
	bufferSize := make([]byte, int(math.Pow(2, 24)))

	// with weighted semaphore by max 10 async goroutines
	ctx := context.Background()
	sem := semaphore.NewWeighted(int64(10))

	// and make the storage and the transport for raw data
	rawData := make([]Raw, 0)
	rawChan := make(chan Raw)
	isFinished := false

	for !isFinished {
		n, readErr := reader.Read(bufferSize)
		if readErr != nil && readErr != io.EOF {
			log.Fatal(readErr)
		}
		if readErr == io.EOF {
			isFinished = true
		}
		if n == 0 {
			break
		}
		rawReadData := bufferSize[:n]
		actualLine := make([]byte, len(rawReadData))
		copy(actualLine, rawReadData)

		sem.Acquire(ctx, 1)
		go processChunk(actualLine, sem, rawChan)
		rawData = append(rawData, <-rawChan)
	}
	sem.Acquire(ctx, 10)
	processRawHotels(rawData)
}

func main() {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		"teknix", "teknixpw", "127.0.0.1", 5432, "hotels")

	pgsql, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Failed to open connection to PostgreSQL: %v", err)
	}

	err = pgsql.Ping()
	if err != nil {
		log.Fatalf("Failed to ping PostgreSQL: %v", err)
	}

	db = pgsql

	fmt.Println("PostgreSQL connection established successfully (for test)")

	parseDump("partner_feed_en_v3_minimal.jsonl.zst")
}

func insertHotel(db *sql.DB, hotel models.Hotel) error {
	// Convert complex fields to JSON
	amenityGroups, err := json.Marshal(hotel.AmenityGroups)
	if err != nil {
		return err
	}
	descriptionStruct, err := json.Marshal(hotel.DescriptionStruct)
	if err != nil {
		return err
	}
	policyStruct, err := json.Marshal(hotel.PolicyStruct)
	if err != nil {
		return err
	}
	roomGroups, err := json.Marshal(hotel.RoomGroups)
	if err != nil {
		return err
	}
	region, err := json.Marshal(hotel.Region)
	if err != nil {
		return err
	}
	serpFilters, err := json.Marshal(hotel.SerpFilters)
	if err != nil {
		return err
	}
	metapolicyStruct, err := json.Marshal(hotel.MetapolicyStruct)
	if err != nil {
		return err
	}
	facts, err := json.Marshal(hotel.Facts)
	if err != nil {
		return err
	}
	paymentMethods, err := json.Marshal(hotel.PaymentMethods)
	if err != nil {
		return err
	}
	images, err := json.Marshal(hotel.Images)
	if err != nil {
		return err
	}

	query := `
	INSERT INTO hotels (
		address, amenity_groups, check_in_time, check_out_time,
		description_struct, id, images, kind, latitude, longitude,
		name, phone, policy_struct, postal_code, room_groups, region,
		star_rating, email, serp_filters, is_closed, is_gender_specification_required,
		metapolicy_struct, metapolicy_extra_info, star_certificate, facts,
		payment_methods, hotel_chain, front_desk_time_start, front_desk_time_end, semantic_version
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30)`

	_, _ = db.Exec(query,
		hotel.Address,                       // $1
		string(amenityGroups),               // $2
		hotel.CheckInTime,                   // $3
		hotel.CheckOutTime,                  // $4
		string(descriptionStruct),           // $5
		hotel.ID,                            // $6
		string(images),                      // $7
		hotel.Kind,                          // $8
		hotel.Latitude,                      // $9
		hotel.Longitude,                     // $10
		hotel.Name,                          // $11
		hotel.Phone,                         // $12
		string(policyStruct),                // $13
		hotel.PostalCode,                    // $14
		string(roomGroups),                  // $15
		string(region),                      // $16
		hotel.StarRating,                    // $17
		hotel.Email,                         // $18
		string(serpFilters),                 // $19
		hotel.IsClosed,                      // $20
		hotel.IsGenderSpecificationRequired, // $21
		string(metapolicyStruct),            // $22
		hotel.MetapolicyExtraInfo,           // $23
		hotel.StarCertificate,               // $24
		string(facts),                       // $25
		string(paymentMethods),              // $26
		hotel.HotelChain,                    // $27
		hotel.FrontDeskTimeStart,            // $28
		hotel.FrontDeskTimeEnd,              // $29
		hotel.SemanticVersion)               // $30

	return nil
}

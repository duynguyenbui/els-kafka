package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"time"

	"github.com/duynguyenbui/async-reader/models"
	"github.com/jackc/pgx/v4"
	"github.com/klauspost/compress/zstd"
)

func main() {
	conn, err := newConn()
	if err != nil {
		fmt.Println("newDB", err)
		return
	}

	defer func() {
		_ = conn.Close(context.Background())
		fmt.Println("closed")
	}()

	now := time.Now()

	if err := insertsHotels(conn); err != nil {
		fmt.Println("failed", err)
		return
	}

	fmt.Println("total", time.Since(now))
}

func insertsHotels(conn *pgx.Conn) error {
	f, err := os.Open("partner_feed_en_v3_minimal.jsonl.zst")
	if err != nil {
		return fmt.Errorf("os.Open %w", err)
	}
	defer f.Close()

	zstdReader, err := zstd.NewReader(f)
	if err != nil {
		return fmt.Errorf("zstd.NewReader %w", err)
	}
	defer zstdReader.Close()

	reader := json.NewDecoder(zstdReader)

	var count int
	for {
		var hotels models.Hotels // Ensure you are using the correct type for Hotels
		if err := reader.Decode(&hotels); err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("json.Decode %w", err)
		}

		amenityGroups, err := json.Marshal(hotels.AmenityGroups)
		if err != nil {
			return fmt.Errorf("json.Marshal amenityGroups %w", err)
		}
		descriptionStruct, err := json.Marshal(hotels.DescriptionStruct)
		if err != nil {
			return fmt.Errorf("json.Marshal descriptionStruct %w", err)
		}
		images, err := json.Marshal(hotels.Images)
		if err != nil {
			return fmt.Errorf("json.Marshal images %w", err)
		}
		policyStruct, err := json.Marshal(hotels.PolicyStruct)
		if err != nil {
			return fmt.Errorf("json.Marshal policyStruct %w", err)
		}
		roomGroups, err := json.Marshal(hotels.RoomGroups)
		if err != nil {
			return fmt.Errorf("json.Marshal roomGroups %w", err)
		}
		region, err := json.Marshal(hotels.Region)
		if err != nil {
			return fmt.Errorf("json.Marshal region %w", err)
		}
		serpFilters, err := json.Marshal(hotels.SerpFilters)
		if err != nil {
			return fmt.Errorf("json.Marshal serpFilters %w", err)
		}
		metapolicyStruct, err := json.Marshal(hotels.MetapolicyStruct)
		if err != nil {
			return fmt.Errorf("json.Marshal metapolicyStruct %w", err)
		}
		metapolicyExtraInfo, err := json.Marshal(hotels.MetapolicyExtraInfo)
		if err != nil {
			return fmt.Errorf("json.Marshal metapolicyExtraInfo %w", err)
		}
		starCertificate, err := json.Marshal(hotels.StarCertificate)
		if err != nil {
			return fmt.Errorf("json.Marshal starCertificate %w", err)
		}
		facts, err := json.Marshal(hotels.Facts)
		if err != nil {
			return fmt.Errorf("json.Marshal facts %w", err)
		}
		paymentMethods, err := json.Marshal(hotels.PaymentMethods)
		if err != nil {
			return fmt.Errorf("json.Marshal paymentMethods %w", err)
		}

		_, err = conn.Exec(context.Background(), `
			INSERT INTO hotels (
				id, address, amenity_groups, check_in_time, check_out_time, description_struct, 
				images, kind, latitude, longitude, name, phone, policy_struct, postal_code, 
				room_groups, region, star_rating, email, serp_filters, is_closed, 
				is_gender_specification_required, metapolicy_struct, metapolicy_extra_info, 
				star_certificate, facts, payment_methods, hotel_chain, front_desk_time_start, 
				front_desk_time_end, semantic_version
			) VALUES (
				$1, $2, $3, $4, $5, $6, 
				$7, $8, $9, $10, $11, $12, $13, $14, 
				$15, $16, $17, $18, $19, $20, 
				$21, $22, $23, $24, $25, $26, $27, 
				$28, $29, $30
			)`,
			hotels.ID, hotels.Address, amenityGroups, hotels.CheckInTime, hotels.CheckOutTime,
			descriptionStruct, images, hotels.Kind, hotels.Latitude, hotels.Longitude,
			hotels.Name, hotels.Phone, policyStruct, hotels.PostalCode, roomGroups,
			region, hotels.StarRating, hotels.Email, serpFilters, hotels.IsClosed,
			hotels.IsGenderSpecificationRequired, metapolicyStruct, metapolicyExtraInfo,
			starCertificate, facts, paymentMethods, hotels.HotelChain, hotels.FrontDeskTimeStart,
			hotels.FrontDeskTimeEnd, hotels.SemanticVersion,
		)
		if err != nil {
			return fmt.Errorf("conn.Exec %w", err)
		}

		count++
	}

	fmt.Println("rows", count)

	return nil
}

func newConn() (*pgx.Conn, error) {
	dsn := url.URL{
		Scheme: "postgres",
		Host:   "localhost:5432",
		User:   url.UserPassword("teknix", "teknixpw"),
		Path:   "hotels",
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")

	dsn.RawQuery = q.Encode()

	conn, err := pgx.Connect(context.Background(), dsn.String())
	if err != nil {
		return nil, fmt.Errorf("pgx.Connect %w", err)
	}

	return conn, nil
}

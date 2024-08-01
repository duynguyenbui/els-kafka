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
	conn, err := NewConnRegions()
	if err != nil {
		fmt.Println("hotels", err)
		return
	}

	defer func() {
		_ = conn.Close(context.Background())
		fmt.Println("closed")
	}()

	now := time.Now()

	if err := insertsRegions(conn); err != nil {
		fmt.Println("failed", err)
		return
	}

	fmt.Println("total", time.Since(now))
}

func insertsRegions(conn *pgx.Conn) error {
	f, err := os.Open("region_v3.jsonl.zst")
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
		var region models.Regions
		if err := reader.Decode(&region); err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("json.Decode %w", err)
		}

		countryName, err := json.Marshal(region.CountryName)
		if err != nil {
			return fmt.Errorf("json.Marshal countryName %w", err)
		}
		center, err := json.Marshal(region.Center)
		if err != nil {
			return fmt.Errorf("json.Marshal center %w", err)
		}
		hotels, err := json.Marshal(region.Hotels)
		if err != nil {
			return fmt.Errorf("json.Marshal hotels %w", err)
		}
		name, err := json.Marshal(region.Name)
		if err != nil {
			return fmt.Errorf("json.Marshal name %w", err)
		}

		_, err = conn.Exec(context.Background(), `
			INSERT INTO regions (
				country_name, country_code, center, hotels, iata, id, type, name
			) VALUES (
				$1, $2, $3, $4, $5, $6, $7, $8
			)`,
			countryName, region.CountryCode, center, hotels, region.Iata, region.ID, region.Type, name,
		)
		if err != nil {
			return fmt.Errorf("conn.Exec %w", err)
		}

		count++
	}

	fmt.Println("rows", count)

	return nil
}

func NewConnRegions() (*pgx.Conn, error) {
	dsn := url.URL{
		Scheme: "postgres",
		Host:   "localhost:5432",
		User:   url.UserPassword("teknix", "teknixpw"),
		Path:   "regions",
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

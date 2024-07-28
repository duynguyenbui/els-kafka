package initialize

import (
	"database/sql"
	"fmt"
	"internacs-els-kafka/global"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type PostgresConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func InitPostgres() {

	port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))

	pgConfig := PostgresConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     port,
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DATABASE"),
	}

	fmt.Printf("Connecting to PostgreSQL with the following configuration: %+v\n", pgConfig)

	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		pgConfig.Username, pgConfig.Password, pgConfig.Host, pgConfig.Port, pgConfig.Database)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Failed to open connection to PostgreSQL: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping PostgreSQL: %v", err)
	}

	global.Pdb = db
	
	fmt.Println("Connected to PostgreSQL successfully!")
}

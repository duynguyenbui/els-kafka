package global

import (
	"database/sql"

	elasticsearchv7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/redis/go-redis/v9"
)

var (
	Rdb *redis.Client
	Els *elasticsearchv7.Client
	Pdb *sql.DB
)

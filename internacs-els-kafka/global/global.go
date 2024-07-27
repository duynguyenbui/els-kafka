package global

import (
	"database/sql"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/redis/go-redis/v9"
)

var (
	Rdb *redis.Client
	Els *elasticsearch.Client
	Pdb *sql.DB
)

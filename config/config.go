package config

import (
	"database/sql"
	"net/http"
)

type CliConfig struct {
	Client *http.Client
	Db     *sql.DB
}

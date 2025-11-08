package cli

import (
	"database/sql"
	"net/http"

	"github.com/ManoloEsS/burrow_prototype/config"
)

type State struct {
	Db     *sql.DB
	Cfg    *config.CliConfig
	Client *http.Client
}

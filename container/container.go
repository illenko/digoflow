package container

import (
	"database/sql"

	"github.com/go-resty/resty/v2"
)

type Container struct {
	Database   *sql.DB
	HttpClient *resty.Client
}

func NewContainer(db *sql.DB, httpClient *resty.Client) *Container {
	return &Container{
		Database:   db,
		HttpClient: httpClient,
	}
}

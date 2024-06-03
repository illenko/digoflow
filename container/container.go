package container

import "database/sql"

type Container struct {
	db *sql.DB
}

func NewContainer(db *sql.DB) *Container {
	return &Container{db: db}
}

func (c *Container) GetDB() *sql.DB {
	return c.db
}

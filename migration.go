package digoflow

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
)

func dbMigration(migrationsDir string, db *sql.DB) (err error) {

	err = goose.SetDialect("postgres")
	if err != nil {
		fmt.Println("When setting database dialect")
		return
	}

	if err != nil {
		return
	}

	err = goose.Up(db, migrationsDir)

	if err != nil {
		fmt.Println("When executing migration")
		return
	}

	return nil
}

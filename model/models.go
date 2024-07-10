package modelofApp

import (
	_ "github.com/mattn/go-sqlite3"
)

type Note struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

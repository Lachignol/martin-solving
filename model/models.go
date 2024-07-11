package modelofApp

import (
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Note struct {
	Id           int        `db:"id"`
	Title        string     `db:"title"`
	Completed    bool       `db:"completed"`
	Created_at   time.Time  `db:"created_at"`
	Completed_at *time.Time `db:"completed_at"`
}

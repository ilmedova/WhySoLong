package models

import "time"

type URL struct {
	ID        string    `db:"id"`
	Original  string    `db:"original"`
	Shortened string    `db:"shortened"`
	CreatedAt time.Time `db:"created_at"`
}

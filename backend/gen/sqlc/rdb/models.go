// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package rdb

import (
	"database/sql"
)

type Task struct {
	ID          uint64
	Title       string
	IsCompleted sql.NullBool
	UpdatedAt   uint64
	CreatedAt   uint64
}

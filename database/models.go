// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Message struct {
	MessageID   int64
	UserID      int32
	MessageText string
	MessageType string
	CreatedAt   pgtype.Timestamptz
}

type User struct {
	UserID    int64
	UserName  string
	UserRole  string
	Email     string
	Password  string
	CreatedAt pgtype.Timestamptz
}

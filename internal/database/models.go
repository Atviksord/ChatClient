// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"time"
)

type Message struct {
	ID      int32
	UserID  int32
	RoomID  int32
	Content string
	SentAt  time.Time
}

type Room struct {
	ID        int32
	Name      string
	CreatedAt time.Time
}

type User struct {
	ID        int32
	Username  string
	Password  string
	Role      sql.NullString
	Status    sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRoom struct {
	UserID   int32
	RoomID   int32
	JoinedAt time.Time
}
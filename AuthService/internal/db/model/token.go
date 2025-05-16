package model

import "time"

type Token struct {
	Id        int64     `db:"id"`
	UserId    int64     `db:"user_id"`
	Token     string    `db:"token"`
	ExpiresAt time.Time `db:"expires_at"`
	Revoked   bool      `db:"reavoked"`
}

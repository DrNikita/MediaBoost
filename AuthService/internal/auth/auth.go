package auth

import "time"

type User struct {
	ID                 int64
	Email              string
	PasswordSaltedHash string
	FirstName          string
	SecondName         string
	Nickname           string
	Age                int
	Vip                bool
	VipExpirationDate  time.Time
}

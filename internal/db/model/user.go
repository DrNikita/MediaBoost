package model

import "time"

type User struct {
	Id                int64     `db:"id"`
	Role              int       `db:"user_role"`
	PasswordHash      string    `db:"password_hash"`
	Email             string    `db:"email"`
	Nickname          string    `db:"nickname"`
	FirstName         string    `db:"first_name"`
	SecondName        string    `db:"second_name"`
	Age               int       `db:"age"`
	ImgPath           string    `db:"img_path"`
	VipExpirationDate time.Time `db:"vip_expiration_date"`
}

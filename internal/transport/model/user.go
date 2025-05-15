package model

type User struct {
	Id         int64  `json:"id"`
	Role       int    `json:"user_role"`
	Age        int    `json:"age"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Nickname   string `json:"nickname"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}

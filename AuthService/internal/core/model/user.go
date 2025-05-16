package model

type User struct {
	Id         int64  `json:"id"`
	Age        int    `json:"age"`
	Role       string `json:"user_role"`
	Email      string `json:"email"`
	Nickname   string `json:"nickname"`
	Password   string `json:"password"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}

func NewUser(age int, role, email, nickname, password, firstName, secondName string) User {
	return User{
		Role:       role,
		Age:        age,
		Password:   password,
		Email:      email,
		Nickname:   nickname,
		FirstName:  firstName,
		SecondName: secondName,
	}
}

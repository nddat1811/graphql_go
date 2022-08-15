// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type LoginResponse struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

type NewUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Respone struct {
	Msg string `json:"msg"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

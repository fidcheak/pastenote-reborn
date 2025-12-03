package main

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
type UserSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

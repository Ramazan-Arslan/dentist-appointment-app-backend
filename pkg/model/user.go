package model

import "errors"

//User data
type User struct {
	Userdata struct {
		ID       int    `json:"id"`
		Fullname string `json:"fullName"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Role     string `json:"role"`
		Password string `json:"password"`
	} `json:"userData"`
	Accesstoken  string `json:"accessToken"`
	Refreshtoken string `json:"refreshToken"`
}

var (
	ErrInvalidUserID = errors.New("invalid user id")
	ErrEmptyUserMail = errors.New("user mail is empty")
)

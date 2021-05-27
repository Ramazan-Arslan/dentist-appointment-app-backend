package model

//Doctor data
type Doctor struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
	Gain     int64  `json:"gain"`
	Phone    string `json:"phone"`
}

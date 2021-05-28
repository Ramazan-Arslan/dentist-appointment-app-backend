package model

//Doctor data
type Type struct {
	ID    uint   `json:"id"`
	Type  string `json:"type"`
	Price int64  `json:"price"`
}

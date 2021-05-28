package model

//Doctor data
type Appointment struct {
	ID            uint    `json:"id"`
	Doctor        *Doctor `json:"doctor"`
	Type          *Type   `json:"type"`
	PatientName   string  `json:"patient_name"`
	PatientAge    int     `json:"patient_age"`
	PatientGender string  `json:"patient_gender"`
	Date          float64 `json:"date"`
	Hour          string  `json:"hour"`
	PatientPhone  string  `json:"patient_phone"`
	Description   string  `json:"description"`
}

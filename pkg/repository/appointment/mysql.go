package appointment

import (
	"database/sql"
	"fmt"

	"github.com/ceng316/dentist-backend/pkg/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

type MySQLRepository struct {
	db *sql.DB
}

const (
	tableName = "appointments"
)

const (
	initTableTemplate = `
	CREATE TABLE IF NOT EXISTS %s (
		id bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
		doctor_id bigint(20) NOT NULL DEFAULT 0,
		type_id bigint(20) NOT NULL DEFAULT 0,
		patient_name varchar(256) NOT NULL DEFAULT '',
		patient_age bigint(20) NOT NULL DEFAULT 0,
		patient_gender varchar(256) NOT NULL DEFAULT '',
		patient_phone varchar(256) NOT NULL DEFAULT '',
		date bigint(20) NOT NULL DEFAULT 0,
		hour varchar(256) NOT NULL DEFAULT '',
		description varchar(256) NOT NULL DEFAULT '',
		UNIQUE KEY id (id)
	  ) ENGINE=MyISAM  DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;	
`
)

func NewMySQLRepository(db *sql.DB) (*MySQLRepository, error) {
	tableInitCmd := fmt.Sprintf(initTableTemplate, tableName)
	_, err := db.Exec(tableInitCmd)

	if err != nil {
		return nil, fmt.Errorf("error init doctor repository: %v", err)
	}

	return &MySQLRepository{
		db: db,
	}, nil
}

// Return all Appointments
func (r *MySQLRepository) GetAll() ([]*model.Appointment, error) {
	q := "SELECT id, doctor_id, type_id, patient_name, patient_age, patient_gender, patient_phone, date, hour, description FROM " + tableName

	res, err := r.db.Query(q)
	if err != nil {
		return nil, err
	}
	var appointments []*model.Appointment

	for res.Next() {
		a := &model.Appointment{}
		d := &model.Doctor{}
		t := &model.Type{}
		a.Type = t
		a.Doctor = d
		err = res.Scan(&a.ID, &a.Doctor.ID, &a.Type.ID, &a.PatientName, &a.PatientAge, &a.PatientGender, &a.PatientPhone, &a.Date, &a.Hour, &a.Description)
		if err != nil {
			logrus.Warn("couldn't scan appointment data from db", err)
			continue
		}
		appointments = append(appointments, a)
	}
	return appointments, nil
}

// Add an Appointments
func (r *MySQLRepository) Add(appointment *model.Appointment) (bool, error) {

	stmt, err := r.db.Prepare("insert into " + tableName + " (doctor_id, type_id, patient_name, patient_age, patient_gender, patient_phone, date, hour, description) VALUES(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return false, err
	}
	_, err = stmt.Exec(appointment.Doctor.ID, appointment.Type.ID, appointment.PatientName, appointment.PatientAge, appointment.PatientGender, appointment.PatientPhone, appointment.Date, appointment.Hour, appointment.Description)
	if err != nil {
		return false, err
	}

	return true, nil
}

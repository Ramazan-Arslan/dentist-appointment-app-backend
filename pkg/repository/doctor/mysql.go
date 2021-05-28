package doctor

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
	tableName = "doctors"
)

const (
	initTableTemplate = `
	CREATE TABLE IF NOT EXISTS %s (
		id bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
		full_name varchar(256) NOT NULL DEFAULT '',
		phone varchar(256) NOT NULL DEFAULT '',
		gain bigint(20) NOT NULL DEFAULT 0,
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

// Return active deviceses count
func (r *MySQLRepository) GetDoctorFromID(id int64) (*model.Doctor, error) {
	q := "SELECT id, full_name, phone, gain  FROM " + tableName + " WHERE id=?"

	logrus.Debug("QUERY: ", q, "id: ", id)
	res := r.db.QueryRow(q, id)

	d := &model.Doctor{}

	if err := res.Scan(&d.ID, &d.FullName, &d.Phone, &d.Gain); err != nil {
		return nil, err
	}

	return d, nil
}

func (r *MySQLRepository) Add(doctor *model.Doctor) (bool, error) {

	stmt, err := r.db.Prepare("insert into " + tableName + " (full_name, phone, gain) VALUES(?,?,?)")
	if err != nil {
		return false, err
	}
	_, err = stmt.Exec(doctor.FullName, doctor.Phone, doctor.Gain)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *MySQLRepository) Update(doctor *model.Doctor) (bool, error) {

	stmt, err := r.db.Prepare("UPDATE " + tableName + " SET full_name=?, phone=?, gain=? WHERE id=?")
	if err != nil {
		return false, err
	}
	_, err = stmt.Exec(doctor.FullName, doctor.Phone, doctor.Gain, doctor.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}
func (r *MySQLRepository) CheckExists(id uint) (bool, error) {
	var exists bool
	row := r.db.QueryRow(`SELECT EXISTS(SELECT 1 FROM `+tableName+` WHERE id=? )`, id)
	if err := row.Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}

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
	tableName = "appointment"
)

const (
	initTableTemplate = `
	CREATE TABLE IF NOT EXISTS %s (
		id bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
		doctor_id varchar(256) NOT NULL DEFAULT 0,
		type_id varchar(256) NOT NULL DEFAULT 0,
		patient_id
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
func (r *MySQLRepository) GetDoctor(user model.Doctor) (*model.Doctor, error) {
	q := "SELECT id, full_name FROM " + tableName + " WHERE id=?"

	logrus.Debug("QUERY: ", q, "id: ", user.ID)
	res := r.db.QueryRow(q, user.ID)

	d := &model.Doctor{}

	if err := res.Scan(&d.ID, &d.FullName); err != nil {
		return nil, err
	}

	return d, nil
}

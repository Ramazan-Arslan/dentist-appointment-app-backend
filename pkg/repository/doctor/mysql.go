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
	q := "SELECT id, full_name FROM " + tableName + " WHERE id=?"

	logrus.Debug("QUERY: ", q, "id: ", id)
	res := r.db.QueryRow(q, id)

	d := &model.Doctor{}

	if err := res.Scan(&d.ID, &d.FullName); err != nil {
		return nil, err
	}

	return d, nil
}

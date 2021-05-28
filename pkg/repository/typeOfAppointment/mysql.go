package typeofappointment

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
	tableName = "type"
)

const (
	initTableTemplate = `
	CREATE TABLE IF NOT EXISTS %s (
		id bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
		type varchar(256) NOT NULL DEFAULT '',
		price bigint(20) NOT NULL DEFAULT 0,
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
func (r *MySQLRepository) GetTypeFromID(id int64) (*model.Type, error) {
	q := "SELECT id, type, price  FROM " + tableName + " WHERE id=?"

	logrus.Debug("QUERY: ", q, "id: ", id)
	res := r.db.QueryRow(q, id)

	t := &model.Type{}

	if err := res.Scan(&t.ID, &t.Type, &t.Price); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return t, nil
}

func (r *MySQLRepository) Add(t *model.Type) (bool, error) {

	stmt, err := r.db.Prepare("insert into " + tableName + " (type, price) VALUES(?,?)")
	if err != nil {
		return false, err
	}
	_, err = stmt.Exec(t.Type, t.Price)
	if err != nil {
		return false, err
	}

	return true, nil
}

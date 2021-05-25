package user

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
	tableName = "users"
)

const (
	initTableTemplate = `
	CREATE TABLE IF NOT EXISTS %s (
		id bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
		email varchar(256) NOT NULL DEFAULT 0,
		password varchar(256) NOT NULL DEFAULT 0,
		refreshToken varchar(256) NOT NULL DEFAULT '',
		token varchar(256) NOT NULL DEFAULT '',
		role varchar(256)  NOT NULL DEFAULT 0,
		fullname varchar(256)  NOT NULL DEFAULT 0,
		UNIQUE KEY id (id)
	  ) ENGINE=MyISAM  DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;	
`
)

func NewMySQLRepository(db *sql.DB) (*MySQLRepository, error) {
	tableInitCmd := fmt.Sprintf(initTableTemplate, tableName)
	_, err := db.Exec(tableInitCmd)

	if err != nil {
		return nil, fmt.Errorf("error init support repository: %v", err)
	}

	return &MySQLRepository{
		db: db,
	}, nil
}

// Return active deviceses count
func (r *MySQLRepository) GetUser(user model.User) (*model.User, error) {
	q := "SELECT email, password, role, fullname, lang FROM " + tableName + " WHERE email=?"

	logrus.Debug("QUERY: ", q, "email: ", user.Userdata.Email)
	res := r.db.QueryRow(q, user.Userdata.Email)

	u := &model.User{}

	if err := res.Scan(&u.Userdata.Email, &u.Userdata.Password, &u.Userdata.Role, &u.Userdata.Fullname, &u.Userdata.Lang); err != nil {
		return nil, err
	}

	return u, nil
}

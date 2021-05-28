package repository

import (
	"database/sql"

	"github.com/ceng316/dentist-backend/pkg/repository/appointment"
	"github.com/ceng316/dentist-backend/pkg/repository/doctor"
	typeofappointment "github.com/ceng316/dentist-backend/pkg/repository/typeOfAppointment"
	"github.com/ceng316/dentist-backend/pkg/repository/user"
)

// MySQL Repository defines the MySQL implementation of Repository interface
type MySQLRepository struct {
	cfg *MySQLConfig
	db  *sql.DB

	userRepo          user.Repository
	doctorRepo        doctor.Repository
	appointmentRepo   appointment.Repository
	typeOfAppointment typeofappointment.Repository
}

// MySQLConfig defines the MySQL Repository configuration
type MySQLConfig struct {
	Addr     string `yaml:"addr"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
}

// dbConn opens connection with MySQL driver
func dbConn(cfg *MySQLConfig) (*sql.DB, error) {

	dbDriver := "mysql"    // Database driver
	dbUser := cfg.Username // Mysql username
	dbPass := cfg.Password // Mysql password
	dbName := cfg.DBName   // Mysql schema
	addr := cfg.Addr

	// Realize the connection with mysql driver
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+addr+")/")
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		return nil, err
	}

	db.Close()

	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
	if err != nil {
		return nil, err
	}

	// Return db object to be used by other functions
	return db, nil
}

// NewMySQLRepository creates a new MySQL Repository
func NewMySQLRepository(cfg *MySQLConfig) (*MySQLRepository, error) {
	db, err := dbConn(cfg)
	if err != nil {
		return nil, err
	}
	userRepo, err := user.NewMySQLRepository(db)
	if err != nil {
		return nil, err
	}
	doctorRepo, err := doctor.NewMySQLRepository(db)
	if err != nil {
		return nil, err
	}
	appointmentRepo, err := appointment.NewMySQLRepository(db)
	if err != nil {
		return nil, err
	}
	typeOfAppointmentRepo, err := typeofappointment.NewMySQLRepository(db)
	if err != nil {
		return nil, err
	}
	return &MySQLRepository{
		cfg:               cfg,
		db:                db,
		userRepo:          userRepo,
		doctorRepo:        doctorRepo,
		appointmentRepo:   appointmentRepo,
		typeOfAppointment: typeOfAppointmentRepo,
	}, nil
}

// GetUserRepository returns the user repository
func (r *MySQLRepository) GetUserRepository() user.Repository {
	return r.userRepo
}

// GetDoctorRepository returns the user repository
func (r *MySQLRepository) GetDoctorRepository() doctor.Repository {
	return r.doctorRepo
}

// GetAppointmentRepository returns the user repository
func (r *MySQLRepository) GetAppointmentRepository() appointment.Repository {
	return r.appointmentRepo
}

// GetAppointmentRepository returns the user repository
func (r *MySQLRepository) GetTypeRepository() typeofappointment.Repository {
	return r.typeOfAppointment
}

// Shutdown closes the database connection
func (r *MySQLRepository) Shutdown() {
	r.db.Close()
}

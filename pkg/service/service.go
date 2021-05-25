package service

import (
	"github.com/ceng316/dentist-backend/pkg/service/user"
)

type Config struct{}

type Service interface {
	GetConfig() *Config
	GetUserService() *user.Service
	Shutdown()
}

package services

import (
	"github.com/deigo96/e-wallet.git/app/repository"
	"github.com/deigo96/e-wallet.git/config"
)

type services struct {
	repositories  repository.Repository
	Configuration *config.Configuration
}

type Services interface{}

func NewService(
	repositories repository.Repository,
	configuration *config.Configuration,
) Services {
	return &services{
		repositories:  repositories,
		Configuration: configuration,
	}
}

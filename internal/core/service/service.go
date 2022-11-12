package service

import (
	"github.com/brightnc/not-human-trading/internal/core/domain"
	"github.com/brightnc/not-human-trading/internal/core/port"
)

/*
	|--------------------------------------------------------------------------
	| Application's Business Logic
	|--------------------------------------------------------------------------
	|
	| Here you can implement a business logic  for your application
	|
*/

type Service struct {
	repository port.Repository
}

func New(repository port.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (svc *Service) SomeBusinessLogic(request domain.BusinessLogicRequest) error {
	return svc.repository.SomeFunction()
}

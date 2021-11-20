package service

import (
	"PersonalCabinetGin/models"
	"PersonalCabinetGin/pkg/repository"
)

type Authorisation interface {

	CreateUser(user models.User)(int, error)
	GenerateToken(username, password string)(string, error)
	ParseToken(token string)(int, error)
}

type Application interface {
	Create(userId int, application models.Application)(int, error)
	GetAll(userId int)( []models.Application, error)
	GetById(userId, applicationId int)(models.Application, error)
	Delete(userId, applicationId int)( error)
}

type Service struct {
	Authorisation
	Application
}

func NewService(repos *repository.Repository) *Service{
	return &Service{
		Authorisation: NewAuthService(repos.Authorisation),
		Application: NewApplicationService(repos.Application),
	}
}
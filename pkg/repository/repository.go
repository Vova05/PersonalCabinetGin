package repository

import (
	"PersonalCabinetGin/models"
	"gorm.io/gorm"
)

type Authorisation interface {
	CreateUser(user models.User)(int,error)
	GetUser(username,password string)(models.User,error)
}

type Application interface {
	Create(userId int, application models.Application)(int, error)
	GetAll(userId int)([]models.Application, error)
	GetById(userId, applicationId int)(models.Application, error)
	Delete(userId, applicationId int)( error)
}

type Repository struct {
	Authorisation
	Application
}

func NewRepository(db *gorm.DB) *Repository{
	return &Repository{
		Authorisation: NewAuthDB(db),
		Application: NewApplicationDB(db),
	}
}

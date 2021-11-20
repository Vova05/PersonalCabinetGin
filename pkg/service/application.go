package service

import (
	"PersonalCabinetGin/models"
	"PersonalCabinetGin/pkg/repository"
)

type ApplicationService struct {
	repo repository.Application
}

func NewApplicationService(repo repository.Application) *ApplicationService{
	return &ApplicationService{repo: repo}
}

func (s *ApplicationService)Create(userId int, application models.Application)(int, error){
	return s.repo.Create(userId, application)
}

func (s *ApplicationService) GetAll(userId int)([]models.Application, error){
	return s.repo.GetAll(userId)
}

func (s *ApplicationService) GetById(userId, applicationId int)(models.Application, error){
	return s.repo.GetById(userId,applicationId)
}

func (s *ApplicationService) Delete(userId, applicationId int)( error){
	return s.repo.Delete(userId,applicationId)
}
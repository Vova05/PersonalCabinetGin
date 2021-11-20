package repository

import (
	"PersonalCabinetGin/models"
	"gorm.io/gorm"
)

type ApplicationDB struct {
	db *gorm.DB
}

func NewApplicationDB(db *gorm.DB) *ApplicationDB{
	return &ApplicationDB{db: db}
}

func (r *ApplicationDB) Create(userId int, application models.Application)(int, error){
	var id int
	//query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", Table)
	application.CreatorId=userId
	row := r.db.Select("Title","Message","CreatorId").Create(&application)
	if err := row.Scan(&application).Error; err != nil {
		return 0,err
	}
	id = application.Id
	return id, nil
}

func (r *ApplicationDB) GetAll(userId int)([]models.Application, error){
	var applications []models.Application
	err := r.db.Table("applications").Where("creator_id = ?", userId).Find(&applications).Error
	return applications,  err
}

func (r *ApplicationDB) GetById(userId, applicationId int)(models.Application, error){
	var application models.Application
	err := r.db.Table("applications").Where("creator_id = ? AND id = ?",userId,applicationId).Scan(&application).Error
	return application, err
}
func (r *ApplicationDB) Delete(userId, applicationId int)( error){
	var app models.Application
	err := r.db.Where("creator_id = ? AND id = ?",userId,applicationId).First(&app).Error
	if err == nil {
		err = r.db.Table("applications").Where("creator_id = ? AND id = ?",userId,applicationId).Delete(app).Error
	}else{
		return err
	}
	return err
}
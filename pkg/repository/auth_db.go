package repository

import (
	"PersonalCabinetGin/models"
	"gorm.io/gorm"
)

type AuthDB struct {
	db *gorm.DB
}

func NewAuthDB(db *gorm.DB) *AuthDB{
	return &AuthDB{db: db}
}

func (r *AuthDB) CreateUser(user models.User)(int,error){
	var id int
	//query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", Table)
	row := r.db.Select("Name","Username","Password").Create(&user)
	if err := row.Scan(&user).Error; err != nil {
		return 0,err
	}
	id = user.Id
	return id, nil
}

func (r *AuthDB) GetUser(username,password string)(models.User,error){
	var user models.User
	 err := r.db.Where("username = ? AND password = ?",username,password).First(&user).Error
	return user, err
}

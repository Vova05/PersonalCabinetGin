package models

type Application struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Message string `json:"message"`
	CreatorId int `json:"creator"`
}

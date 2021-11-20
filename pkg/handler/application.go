package handler

import (
	"PersonalCabinetGin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createApplication(c *gin.Context){
	userId, err := getUserId(c)

	if err != nil{
		return
	}
	var input models.Application
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}

	id, err := h.service.Application.Create(userId,input)

	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}


	c.JSON(http.StatusOK,map[string]interface{}{
		"id": id,
	})
}

type getAllApplications struct {
	Data []models.Application `json:"data"`
}

func (h *Handler)  getAllApplication(c *gin.Context){

	userId, err := getUserId(c)

	if err != nil{
		return
	}

	applications, err := h.service.Application.GetAll(userId)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,getAllApplications{
		Data: applications,
	})
}

func (h *Handler) getApplicationById(c *gin.Context){

	userId, err := getUserId(c)

	if err != nil{
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c,http.StatusBadRequest, "invalid id param")
		return
	}

	application, err := h.service.Application.GetById(userId, id)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,application)
}

func (h *Handler) updateApplication(c *gin.Context){

}

func (h *Handler) deleteApplication(c *gin.Context) {

	userId, err := getUserId(c)

	if err != nil{
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c,http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.service.Application.Delete(userId, id)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Deleted",
	})
}
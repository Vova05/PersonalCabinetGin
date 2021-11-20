package handler

import (
	"PersonalCabinetGin/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler{
	return &Handler{service: services}
}

func (h *Handler) InitRoutes() *gin.Engine{
	router := gin.New()

	router.Static("/css","C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin\\pkg\\handler\\templates\\css")
	router.Static("/fonts","C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin\\pkg\\handler\\templates/css/fonts")
	router.Static("/images","C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin\\pkg\\handler\\templates/images")
	router.Static("/js","C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin\\pkg\\handler\\templates/js")

	router.LoadHTMLGlob("C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin\\pkg\\handler\\templates/*.html")

	auth := router.Group("/auth")
	{
		auth.Static("/css","C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin\\pkg\\handler\\templates\\css")
		auth.Static("/fonts","C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin\\pkg\\handler\\templates/assets/css/fonts")
		auth.Static("/images","C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin\\pkg\\handler\\templates/assets/images")
		auth.Static("/js","C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin\\pkg\\handler\\templates/js")

		auth.POST("/sign-up", h.signUp)

		auth.POST("/sign-in", h.signIn)
		auth.GET("/sign-in",h.signInGet)
	}

	api :=router.Group("/api",h.userIdentity)
	{
		application := api.Group("/applications")
		{
			application.GET("/", h.getAllApplication)
			application.GET("/:id",h.getApplicationById)
			application.POST("/create", h.createApplication)
			//application.PUT("/:id", h.updateApplication)
			application.DELETE("/:id",h.deleteApplication)
		}
	}
	return router
}

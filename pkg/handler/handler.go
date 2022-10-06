package handler

import (
	_ "github.com/Mirobidjon/contact-list/docs"
	"github.com/Mirobidjon/contact-list/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Handler struct
type Handler struct {
	service *service.Service
}

// NewHandler ...
func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// InitRoutes  ...
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentify)
	{
		contacts := api.Group("/contact")
		{
			contacts.GET("/", h.getAllContact)
			contacts.POST("/", h.createContact)
			contacts.PUT("/:id", h.updateContact)
			contacts.DELETE("/:id", h.deleteContact)
			contacts.GET("/:id", h.getContact)
		}
	}

	// swagger
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

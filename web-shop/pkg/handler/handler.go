package handler

import (
	"github.com/evgenchikk/itmo-web-shop/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	api := router.Group("api")
	{
		catalog := api.Group("catalog")
		{
			catalog.GET("groups", h.getAllItemGroups)
			catalog.GET("groups/:id", h.getItemGroupById)
			catalog.GET("groups/:id/items", h.getItemsByItemGroupId)

			catalog.GET("items", h.getAllItems)
			catalog.GET("items/:id", h.getItemById)

			catalog.GET("categories", h.getAllCategories)
			catalog.GET("categories/:id", h.getCategoryById)
			catalog.GET("categories/:id/groups", h.getItemGroupsByCategoryId)
		}

	}

	return router
}

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

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	api := router.Group("api")
	{
		api.GET("item_groups", h.getAllItemGroups)
		api.GET("item_groups/:id", h.getItemGroupById)
		api.GET("item_groups/:id/items", h.getItemsByItemGroupId)

		api.GET("items", h.getAllItems)
		api.GET("items/:id", h.getItemById)

		api.GET("users")
	}

	return router
}

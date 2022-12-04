package handler

import (
	"net/http"

	"github.com/evgenchikk/itmo-web-shop/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var requestJSON models.User

	if err := c.BindJSON(&requestJSON); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error msg": err.Error()})
		return
	}

	if id, err := h.services.Authorization.SignUp(requestJSON); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error msg": err.Error()})
	} else {
		c.IndentedJSON(http.StatusCreated, gin.H{"id": id})
	}
}

func (h *Handler) signIn(c *gin.Context) {

}

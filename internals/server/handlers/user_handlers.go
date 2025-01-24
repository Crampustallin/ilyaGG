package handlers

import (
	"net/http"

	"github.com/Crampustallin/ilyaGG/internals/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetUserHandlersGroup(appApi *gin.RouterGroup) {
	userGroup := appApi.Group("/user")
	{
		userGroup.POST("/register", h.register)
		userGroup.GET("/check", h.check)
	}
}

func (h *Handler) register(c *gin.Context) {
	var user *models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
}

func (h *Handler) check(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ok": "hello world"})
}

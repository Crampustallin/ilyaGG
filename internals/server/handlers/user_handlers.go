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
		userGroup.POST("/login", h.login)
		userGroup.GET("/check", h.check)
	}
}

func (h *Handler) register(c *gin.Context) {
	var u *models.UserRegister
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	if err := h.service.UserService.SetUser(u.UserName, u.UserPass); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"ok": u})
}

func (h *Handler) login(c *gin.Context) {
	var user *models.UserRegister
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	if err := h.service.UserService.GetUser(user.UserName, user.UserPass); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": "ok"})
}

func (h *Handler) check(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ok": "hello world"})
}

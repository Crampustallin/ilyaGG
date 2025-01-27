package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/Crampustallin/ilyaGG/views"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

const appTimeout = time.Second * 10

func render(c *gin.Context, status int, template templ.Component) {
	c.Status(status)
	template.Render(c.Request.Context(), c.Writer)
}

func (h *Handler) IndexHandler(c *gin.Context) {
	_, cancel := context.WithTimeout(context.Background(), appTimeout)
	defer cancel()
	render(c, http.StatusOK, views.Index())
}

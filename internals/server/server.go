package server

import (
	"net/http"
	"os"

	"github.com/Crampustallin/ilyaGG/internals/server/handlers"
	"github.com/gin-gonic/gin"
)

const SERVER_ADDR = "SERVER_ADDR"

func NewServer() *http.Server {
	router := gin.Default()
	router.Use(corsMiddleWare())
	appApi := router.Group("app-api")

	handler := handlers.New()

	handler.SetUserHandlersGroup(appApi)

	return &http.Server{
		Addr:    os.Getenv(SERVER_ADDR),
		Handler: router,
	}
}

func corsMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()

	}
}

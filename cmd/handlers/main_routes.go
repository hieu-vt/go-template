package handlers

import (
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"go-template/common"
	"net/http"
)

func MainRoute(router *gin.Engine, sc goservice.ServiceContext) {
	authClient := sc.MustGet(common.PluginGrpcAuthClient).(interface {
		RequiredAuth(sc goservice.ServiceContext) func(c *gin.Context)
	})

	// use authClient.RequiredAuth(sc)

	v1 := router.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "pong go template"})
		})

		v1.GET("/auth/ping", authClient.RequiredAuth(sc), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "Auth pong"})
		})
	}
}

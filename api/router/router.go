package router

import (
	"github.com/anneau/go-template/api/controller"
	"github.com/gin-gonic/gin"
)

type Router struct {
	HealthController *controller.HealthController
}

func (r *Router) Setup(route *gin.RouterGroup)  {
	route.GET("/health", r.HealthController.Check)
}

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

type HealthDTO struct {
	Status string `json:"status"`
}

func (c *HealthController) Check(ctx *gin.Context)  {
	res := HealthDTO{
		Status: "ok",
	}
	ctx.JSON(http.StatusOK, res)
}

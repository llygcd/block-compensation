package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/llygcd/block-compensation/internal/service"
	"strconv"
)

type CompensationController struct {
	compensationService *service.CompensationService
}

func NewCompensationController(compensationService *service.CompensationService) *CompensationController {
	return &CompensationController{compensationService: compensationService}
}

func (ctl *CompensationController) Compensation(c *gin.Context) {
	var locality = c.Param("height")

	height, err := strconv.ParseInt(locality, 10, 64)
	if err != nil {

	}
	ctl.compensationService.Compensation(height)
}

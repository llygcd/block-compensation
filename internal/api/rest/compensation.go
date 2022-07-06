package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/llygcd/block-compensation/internal/service"
	"github.com/llygcd/block-compensation/pkg/pool"
	"strconv"
)

type CompensationController struct {
	compensationService *service.CompensationService
	poolClient          *pool.Client
}

func NewCompensationController(compensationService *service.CompensationService, poolClient *pool.Client) *CompensationController {
	return &CompensationController{compensationService: compensationService, poolClient: poolClient}
}

func (ctl *CompensationController) Compensation(c *gin.Context) {
	var locality = c.Param("height")

	height, err := strconv.ParseInt(locality, 10, 64)
	if err != nil {

	}
	ctl.compensationService.Compensation(height, ctl.poolClient)

}

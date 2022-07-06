package api

import (
	"github.com/gin-gonic/gin"
	"github.com/llygcd/block-compensation/internal/api/rest"
)

func InitCompensationRouter(r *gin.RouterGroup, ctl *rest.CompensationController) {
	compensationRouter := r.Group("/compensation")

	{
		compensationRouter.GET("/:height", ctl.Compensation)
	}
}

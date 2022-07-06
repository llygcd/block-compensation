package initialization

import (
	"github.com/gin-gonic/gin"
	"github.com/llygcd/block-compensation/internal/api"
	"github.com/llygcd/block-compensation/internal/model"
	"github.com/sirupsen/logrus"
)

func Routers(Router *gin.Engine, ctl *model.Controllers) {

	commonGroup := Router.Group("")

	// use auth for other Router
	commonGroup.Use()
	{

		api.InitCompensationRouter(commonGroup, ctl.CompensationController)

	}
	logrus.Debug("router register success")
}

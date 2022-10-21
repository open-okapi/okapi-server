package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/open-okapi/okapi-server/config"
	"github.com/open-okapi/okapi-server/inner/middleware"
	"github.com/open-okapi/okapi-server/inner/pkg/errors"
	response2 "github.com/open-okapi/okapi-server/inner/pkg/response"
	"io"
	"net/http"
)

func InitRouters() *gin.Engine {
	var r *gin.Engine
	if config.Config.Debug == false {
		// prod
		r = ReleaseRouter()
		r.Use(
			middleware.RequestCostHandler(),
			middleware.CustomLogger(),
			middleware.CustomRecovery(),
			middleware.CorsHandler(),
		)
	} else {
		// dev
		r = gin.New()
		r.Use(
			middleware.RequestCostHandler(),
			gin.Logger(),
			middleware.CustomRecovery(),
			middleware.CorsHandler(),
		)
	}
	// add health check api
	r.Any("/ok", func(context *gin.Context) {
		context.AbortWithStatusJSON(http.StatusOK, gin.H{"health": "ok"})
	})
	// add api
	setApiRouter(r)
	// not found router
	r.NoRoute(func(c *gin.Context) {
		response2.Resp().SetHttpCode(http.StatusNotFound).FailCode(c, errors.NotFound)
	})
	return r
}

func ReleaseRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	engine := gin.New()
	return engine
}

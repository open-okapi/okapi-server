package boot

import (
	"fmt"
	"github.com/open-okapi/okapi-server/config"
	"github.com/open-okapi/okapi-server/data"
	"github.com/open-okapi/okapi-server/inner/pkg/logger"
	"github.com/open-okapi/okapi-server/inner/routers"
	"github.com/open-okapi/okapi-server/inner/validator"
)

func init() {
	// version
	println("current server version : " + version)
	// init config
	config.InitConfig()
	// init logger
	logger.InitLogger()
	// init database
	data.InitData()
	// init validation
	validator.InitValidatorTrans("zh")
}

func Run() {
	r := routers.InitRouters()
	err := r.Run(fmt.Sprintf("%s:%d", config.Config.Server.Host, config.Config.Server.Port))
	if err != nil {
		panic(err)
	}
}

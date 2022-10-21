package config

import (
	. "github.com/open-okapi/okapi-server/config/autoload"
	"github.com/open-okapi/okapi-server/pkg/utils"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

// Conf 配置项主结构体
type Conf struct {
	AppConfig `ini:"app" yaml:"app"`
	Server    ServerConfig `ini:"server" yaml:"server"`
	Mysql     MysqlConfig  `ini:"mysql" yaml:"mysql"`
	Redis     RedisConfig  `ini:"redis" yaml:"redis"`
	Logger    LoggerConfig `ini:"logger" yaml:"logger"`
}

var Config = &Conf{
	AppConfig: App,
	Server:    Server,
	Mysql:     Mysql,
	Redis:     Redis,
	Logger:    Logger,
}

var once sync.Once

func InitConfig() {
	println("server env : " + Config.AppConfig.AppEnv)
	once.Do(func() {
		// 加载 .yaml 配置
		loadYaml()
	})
	println("server env : " + Config.AppConfig.AppEnv)
	println("server debug : " + strconv.FormatBool(Config.AppConfig.Debug))
	println("server language : " + Config.AppConfig.Language)
}

func loadYaml() {
	runDirectory, _ := utils.GetDefaultPath()
	println("init config path : " + runDirectory)
	// 获取 config.yaml 文件路径
	var yamlConfig = filepath.Join(runDirectory, "/config.yaml")
	cfg, err := os.ReadFile(yamlConfig)
	if err != nil {
		panic("Failed to read configuration file:" + err.Error())
	}
	err = yaml.Unmarshal(cfg, &Config)
	if err != nil {
		panic("Failed to load configuration:" + err.Error())
	}
}

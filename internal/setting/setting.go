package setting

import (
	"io/ioutil"
	"log"
	"path"
	"runtime"
	"time"

	"gopkg.in/yaml.v2"
)

// AppSection app.yaml server 配置
type AppSection struct {
	JwtSecret string `yaml:"jwt_secret"`
}

// ServerSection app.yaml server 配置
type ServerSection struct {
	HTTPPort int `yaml:"http_port"`
}

// DatabaseSection app.yaml database 配置
type DatabaseSection struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// RedisSection app.yaml redis 配置
type RedisSection struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// EmailSection app.yaml email 配置
type EmailSection struct {
	From        string `yaml:"from"`
	FromSubject string `yaml:"from_subject"`
	Subject     string `yaml:"subject"`
}

// CodeSection app.yaml code 配置
type CodeSection struct {
	Min            int           `yaml:"min"`
	Max            int           `yaml:"max"`
	ValidityPeriod time.Duration `yaml:"validity_period"`
}

// 所有的配置项
var (
	App      AppSection
	Server   ServerSection
	Database DatabaseSection
	Redis    RedisSection
	Email    EmailSection
	Code     CodeSection
)

// app.yaml 对应结构体
var appConfig struct {
	App      AppSection
	Server   ServerSection
	Database DatabaseSection
	Redis    RedisSection
	Email    EmailSection
	Code     CodeSection
}

// init 初始化加载配置文件
func init() {
	// 解析 app.yml
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("获取文件路径")
	}
	filepath := path.Join(path.Dir(filename), "../../configs/app.yml")
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal([]byte(file), &appConfig)
	if err != nil {
		log.Fatalln(err)
	}
	App = appConfig.App
	Server = appConfig.Server
	Database = appConfig.Database
	Redis = appConfig.Redis
	Email = appConfig.Email
	Code = appConfig.Code
}

package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/li1553770945/openmcp-gateway/biz/constant"
	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	HttpServerListenAddress string `yaml:"listen-address"`
}

type DatabaseConfig struct {
	Type     string `yaml:"type"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Address  string `yaml:"address"`
	Port     int32  `yaml:"port"`
	UseTLS   bool   `yaml:"use-tls"`
}

type AuthConfig struct {
	JWTKey string `yaml:"jwt-key"`
}

type ProxyConfig struct {
	// 缓存过期时间
	CacheExpirationSeconds int32 `yaml:"cache-expiration-seconds"`
}
type Config struct {
	Env            string         `yaml:"-"`
	ServerConfig   ServerConfig   `yaml:"server"`
	DatabaseConfig DatabaseConfig `yaml:"database"`
	AuthConfig     AuthConfig     `yaml:"auth"`
	ProxyConfig    ProxyConfig    `yaml:"proxy"`
}

func GetConfig(env string) *Config {
	if env != constant.EnvProduction && env != constant.EnvDevelopment {
		panic(fmt.Sprintf("环境必须是%s或者%s之一", constant.EnvProduction, constant.EnvDevelopment))
	}
	conf := &Config{}
	path := filepath.Join("conf", fmt.Sprintf("%s.yml", env))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 进入这里说明文件【不存在】
		panic(fmt.Sprintf("配置文件%s不存在！", path))
	}
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	err = yaml.NewDecoder(f).Decode(conf)
	conf.Env = env
	if err != nil {
		panic(err)
	}

	return conf
}

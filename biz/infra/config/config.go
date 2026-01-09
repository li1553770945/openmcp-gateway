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
		fmt.Printf("警告: 环境 '%s' 未知，自动切换为 '%s'\n", env, constant.EnvDevelopment)
		env = constant.EnvDevelopment
	}
	conf := &Config{}
	path := filepath.Join("conf", fmt.Sprintf("%s.yml", env))

	// 尝试读取配置文件
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("提示: 配置文件 '%s' 不存在，将使用内置默认配置。\n", path)
		return getDefaultConfig(env)
	}

	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("警告: 无法打开配置文件 '%s' (%v)，将使用内置默认配置。\n", path, err)
		return getDefaultConfig(env)
	}
	defer f.Close()

	if err = yaml.NewDecoder(f).Decode(conf); err != nil {
		fmt.Printf("警告: 解析配置文件失败 (%v)，将使用内置默认配置。\n", err)
		return getDefaultConfig(env)
	}

	conf.Env = env
	return conf
}

func getDefaultConfig(env string) *Config {
	return &Config{
		Env: env,
		ServerConfig: ServerConfig{
			HttpServerListenAddress: "0.0.0.0:9000",
		},
		DatabaseConfig: DatabaseConfig{
			Type:     "sqlite",
			Database: "openmcp.db",
		},
		AuthConfig: AuthConfig{
			JWTKey: "openmcp-gateway-default-secret-key",
		},
		ProxyConfig: ProxyConfig{
			CacheExpirationSeconds: 60,
		},
	}
}

package database

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/li1553770945/openmcp-gateway/biz/infra/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabase(conf *config.Config) *gorm.DB {
	dbConfig := conf.DatabaseConfig
	var dialector gorm.Dialector

	switch strings.ToLower(dbConfig.Type) {
	case "sqlite", "sqlite3":
		dialector = sqlite.Open(dbConfig.Database)
	case "mysql", "":
		// 默认为 mysql
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=%s", dbConfig.Username, dbConfig.Password, dbConfig.Address, dbConfig.Port, dbConfig.Database, strconv.FormatBool(dbConfig.UseTLS))
		dialector = mysql.Open(dsn)
	default:
		panic(fmt.Sprintf("不支持的数据库类型: %s", dbConfig.Type))
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic("数据库连接失败:" + err.Error())
	}
	return db
}

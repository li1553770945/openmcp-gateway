package database

import (
	"fmt"
	"strconv"

	"github.com/li1553770945/openmcp-gateway/biz/infra/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(conf *config.Config) *gorm.DB {
	dbConfig := conf.DatabaseConfig

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=%s", dbConfig.Username, dbConfig.Password, dbConfig.Address, dbConfig.Port, dbConfig.Database, strconv.FormatBool(dbConfig.UseTLS))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败:" + err.Error())
	}
	return db
}

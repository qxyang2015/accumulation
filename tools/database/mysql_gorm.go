package database

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type GormDBInfo struct {
	*gorm.DB
	DbType       string
	DbHost       string
	DbPort       string
	DbUserName   string
	DbPassWord   string
	DbName       string
	DbDataSource string
}

func CreateGormDBByYaml(viper *viper.Viper, asectionname string) *GormDBInfo {

	dbInfo := &GormDBInfo{}
	dbInfo.DbType = viper.GetString(asectionname + ".dbtype")
	dbInfo.DbHost = viper.GetString(asectionname + ".dbhost")
	dbInfo.DbName = viper.GetString(asectionname + ".dbname")
	dbInfo.DbUserName = viper.GetString(asectionname + ".dbusername")
	dbInfo.DbPassWord = viper.GetString(asectionname + ".dbpassword")
	dbInfo.DbPort = viper.GetString(asectionname + ".dbport")

	return dbInfo
}

func InitGormMysql(dbInfo *GormDBInfo) (*GormDBInfo, error) {

	dbInfo.DbDataSource = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		dbInfo.DbUserName, dbInfo.DbPassWord, dbInfo.DbHost, dbInfo.DbPort, dbInfo.DbName)

	DB, err := gorm.Open("mysql", dbInfo.DbDataSource)
	if err != nil {
		panic(err)
	}
	dbInfo.DB = DB
	//日志信息、标名是否复数、空闲时最大连接数、最大连接数
	dbInfo.DB.LogMode(false)
	dbInfo.DB.SingularTable(true)
	// 空闲时最大连接数
	dbInfo.DB.DB().SetMaxIdleConns(10)
	// 最大连接数
	dbInfo.DB.DB().SetMaxOpenConns(100)

	dbInfo.DB.DB().SetConnMaxLifetime(60 * time.Second)

	return dbInfo, err
}

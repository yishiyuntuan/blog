package db

import (
	"blog/logger"
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// var DB *gorm.DB
//
// func init() {
// 	DB = GetDB()
// }

// Setup 初始化连接
func GetDB() *gorm.DB {

	// db = newConnection()
	var dbURI string
	var dialector gorm.Dialector

	DatabaseSetting := config().Database
	logger.Log.Debug(DatabaseSetting)
	switch DatabaseSetting.Type {
	case "mysql":
		dbURI = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
			DatabaseSetting.User,
			DatabaseSetting.Password,
			DatabaseSetting.Host,
			DatabaseSetting.Port,
			DatabaseSetting.Name)
		dialector = mysql.New(mysql.Config{
			DSN:                       dbURI, // data source name
			DefaultStringSize:         256,   // default size for string fields
			DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
			DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
			DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
			SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
		})
	case "postgres":
		dbURI = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			DatabaseSetting.Host,
			DatabaseSetting.Port,
			DatabaseSetting.User,
			DatabaseSetting.Name,
			DatabaseSetting.Password)
		dialector = postgres.New(postgres.Config{
			DSN:                  "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		})
	case "sqlite3":
		dbURI = fmt.Sprintf("test.db")
		dialector = sqlite.Open("test.db")
	}
	logger.Log.Debug(dialector)

	conn, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Print(err.Error())
	}
	sqlDB, err := conn.DB()
	if err != nil {
		logger.Log.Error("connect db server failed.")
	}
	sqlDB.SetMaxIdleConns(10)  // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxOpenConns(100) // SetMaxOpenConns sets the maximum number of open connections to the database.
	// 5秒内连接没有活跃的话则自动关闭连接
	sqlDB.SetConnMaxLifetime(time.Minute * 30)
	return conn
}

// Config 配置对象
type Config struct {
	Database *Database `json:"database"`
}

// Setup 配置
func config() *Config {
	// GlobalConfigSetting 配置实例
	var ConfigSetting = &Config{}
	filePtr, err := os.Open("db/config.yaml") // config的文件目录
	if err != nil {
		return nil
	}
	defer filePtr.Close()
	// 创建yaml解码器
	decoder := yaml.NewDecoder(filePtr)
	err = decoder.Decode(ConfigSetting)
	return ConfigSetting
}

// Database 数据库配置对象
type Database struct {
	Type        string `yaml:"type"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	Name        string `yaml:"name"`
	TablePrefix string `yaml:"table_prefix"`
}

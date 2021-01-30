package mysql

import (
	"fmt"
	"github-muzilong/go_iris_admin/bootstrap/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// DB gorm.DB 对象
var DB *gorm.DB

func SetupDB() {
	var err error

	// 初始化 MySQL 连接信息
	var (
		host     = config.GetString("database.mysql.host")
		port     = config.GetString("database.mysql.port")
		database = config.GetString("database.mysql.database")
		username = config.GetString("database.mysql.username")
		password = config.GetString("database.mysql.password")
		charset  = config.GetString("database.mysql.charset")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		username, password, host, port, database, charset, true, "Local")

	// 准备数据库连接池
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("数据库连接错误")
	}

	// 命令行打印数据库请求的信息
	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Println(err)
		panic("数据库连接池错误")
	}

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)

}

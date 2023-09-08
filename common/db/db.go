package db

import (
	"chatgpt-web-new-go/common/config"
	"chatgpt-web-new-go/dao"
	"database/sql"
	"fmt"
	"github.com/mattn/go-sqlite3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var dbTypeInitializer = map[string]func(){
	"mysql":  initMysql,
	"sqlite": initSqlite,
}

func Init() {
	dbType := config.Config.Db.Type
	dbInitializer := dbTypeInitializer[dbType]
	dbInitializer()

	// gorm gen init
	dao.SetDefault(config.DB)
}

func initMysql() {
	dbConfig := config.Config.Db
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Name)

	var err error
	config.DB, err = gorm.Open(
		mysql.New(mysql.Config{
			DSN:                       dsn,   // data source name
			DefaultStringSize:         256,   // default size for string fields
			DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
			DontSupportRenameIndex:    true,  // drop & create when rename messageDao, rename messageDao not supported before MySQL 5.7, MariaDB
			DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
			SkipInitializeWithVersion: false, // autoconfigure based on currently MySQL version
		}),
		&gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold: time.Second,
					LogLevel:      logger.Info,
					Colorful:      true,
				},
			),
		})
	if err != nil {
		panic(err)
	}
	sqlDB, err := config.DB.DB()
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// migrate
	//err = config.DB.AutoMigrate(&user.User{})
	//if err != nil {
	//	panic(err)
	//}
}

// initSqlite 数据库初始化，包括新建数据库（如果还没有建立），基本数据的读写
func initSqlite() {
	sql.Register("sqlite3_simple",
		&sqlite3.SQLiteDriver{
			Extensions: []string{
				"libsimple-osx-x64/libsimple",
			},
		},
	)

	var err error
	config.DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

}

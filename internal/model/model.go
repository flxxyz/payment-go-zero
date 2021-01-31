package model

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"payment/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	gormDB        *gorm.DB
	gormCfg       *gorm.Config
	gormDialector gorm.Dialector
	once          sync.Once
)

// Init orm初始化
func Init(cfg *config.Db) (sqlDB *sql.DB, err error) {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
			cfg.Username,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.Database,
			cfg.Charset,
		)

		gormDialector = mysql.New(mysql.Config{
			DSN:                       dsn,
			DefaultStringSize:         256,
			SkipInitializeWithVersion: false,
		})

		gormCfg = &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   cfg.Prefix,
				SingularTable: true, // 不使用复数表名
			},
		}

		gormDB, err = gorm.Open(gormDialector, gormCfg)
		if err != nil {
			log.Fatal(err)
		}

		sqlDB, err = gormDB.DB()
		if err != nil {
			log.Fatal(err)
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		//sqlDB.SetConnMaxLifetime(time.Hour)
	})

	return
}

// Orm gorm实例
func Orm() *gorm.DB {
	return gormDB
}

// Prefix 获取表前缀
func Prefix(name string) string {
	return gormDB.NamingStrategy.TableName(name)
}

package dao

import (
	"database/sql"
	"fmt"
	"gin-vben-admin/global"
	"gin-vben-admin/pkg/logger"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func Init() *gorm.DB {
	logrus.Infof("初始化数据库连接")
	var err error
	var db *gorm.DB
	switch global.Conf.System.DbType {
	case "sqlite", "sqlite3", "":
		db, err = gorm.Open(sqlite.Open(global.Conf.Sqlite.Dsn()))
	case "pgsql":
		c := global.Conf.Pgsql
		if db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  c.Dsn(), // DSN data source name
			PreferSimpleProtocol: false,
		}), &gorm.Config{
			SkipDefaultTransaction:                   true,
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: c.Prefix,
			},
			Logger: logger.New(),
		}); err != nil {
			panic(err)
		}
	case "mysql":
		// 当前只支持 sqlite3 与 mysql 数据库
		c := global.Conf.Mysql
		dsn := c.Dsn()
		createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", c.Dbname)
		if err := createDatabase(dsn, "mysql", createSql); err != nil {
			logrus.WithError(err).Info("mysql db创建失败")
			panic(err)
		} // 创建数据库

		if db, err = gorm.Open(mysql.New(mysql.Config{
			DSN:                       c.Dsn(), // DSN data source name
			SkipInitializeWithVersion: true,    // 根据版本自动配置
		}), &gorm.Config{
			SkipDefaultTransaction:                   true,
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: c.Prefix,
			},
			Logger: logger.New(),
		}); err != nil {
			logrus.WithError(err).Info("mysql连接失败")
		} else {
			dbi, _ := db.DB()
			//空闲
			dbi.SetMaxIdleConns(c.MaxIdleConns)
			//打开
			dbi.SetMaxOpenConns(c.MaxOpenConns)
		}
	}

	if err != nil {
		logrus.Fatalf("连接数据库不成功, %s", err)
	}

	global.DB = db
	return db
}

// createDatabase 创建数据库（ EnsureDB() 中调用 ）
func createDatabase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

func TableName(name string) string {
	return global.Conf.Mysql.Prefix + name
}

func Now() *time.Time {
	t := time.Now()
	return &t
}

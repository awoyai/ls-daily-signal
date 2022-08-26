package xgorm

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sync"
	"time"
)

type XGorm struct {
	*gorm.DB
	sqlDB  *sql.DB
	config *Config
}

// Config 数据库配置
type Config struct {
	// 数据库类型
	DBType DBType
	// 用户名
	DBUser string
	// 密码
	DBPassword string
	// db url
	DBUrl string
	// port
	Port int
	// 数据库名称
	Database string
	// DryRun generate sql without execute
	DryRun bool
	// the maximum number of connections in the idle connection pool
	MaxIdleConns int
	// the maximum number of open connections to the database
	MaxOpenConns int
	// the maximum amount of time a connection may be reused
	ConnMaxLifeTime string
	SSLMode         string
}

// DBType 数据库类型
type DBType string

const (
	DBTypeMySQL    DBType = "mysql"
	DBTypePostgres DBType = "postgres"
)

func (conf *Config) Dialector(initDB ...bool) gorm.Dialector {
	switch conf.DBType {
	case DBTypeMySQL:
		if len(initDB) > 0 && initDB[0] {
			return mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/",
				conf.DBUser, conf.DBPassword, conf.DBUrl))
		}
		// mysql: url 示例: 127.0.0.1:3306(包含端口)
		return mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			conf.DBUser, conf.DBPassword, conf.DBUrl, conf.Database))
	case DBTypePostgres:
		// postgres: url 示例: 1270.0.0.1(不包含端口), port:3306
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Shanghai",
			conf.DBUrl, conf.DBUser, conf.DBPassword, conf.Database, conf.Port, conf.SSLMode)
		return postgres.Open(dsn)
	default:
		panic("unknown DBType")
	}
}

// OpenDB
func OpenDB(config *Config, initDB ...bool) (*XGorm, func(), error) {
	db, err := gorm.Open(config.Dialector(initDB...), &gorm.Config{
		DryRun:  config.DryRun,
		Plugins: nil,
	})
	if err != nil {
		return nil, nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}
	cleanUp := func() {
		_ = sqlDB.Close()
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, cleanUp, err
	}
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	dur, err := time.ParseDuration(config.ConnMaxLifeTime)
	if err != nil {
		return nil, cleanUp, err
	}
	sqlDB.SetConnMaxLifetime(dur)
	return &XGorm{
		DB:     db,
		sqlDB:  sqlDB,
		config: config,
	}, cleanUp, nil
}

// OpenMigrateDB 用来创建数据库
func OpenMigrateDB(config *Config) (*XGorm, func(), error) {
	return OpenDB(config, true)
}

func (x *XGorm) CreateDB() {
	createDbSQL := "CREATE DATABASE IF NOT EXISTS " + x.config.Database + " DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;"
	err := x.Exec(createDbSQL).Error
	if err != nil {
		panic("创建失败：" + err.Error() + " sql:" + createDbSQL)
	}
	fmt.Println(fmt.Sprintf("database: %v create success", x.config.Database))
}

func (x *XGorm) DropDB() {
	dropDbSQL := "DROP DATABASE IF EXISTS " + x.config.Database + ";"
	err := x.Exec(dropDbSQL).Error
	if err != nil {
		panic("删除失败：" + err.Error() + " sql:" + dropDbSQL)
	}
	fmt.Println(fmt.Sprintf("database: %v drop success", x.config.Database))
}

func (x *XGorm) ClearTables(models ...interface{}) error {
	for _, tb := range models {
		sch, err := schema.Parse(tb, &sync.Map{}, x.DB.NamingStrategy)
		if err != nil {
			return errors.New("get table name fail:" + err.Error())
		}
		if err := x.DB.Exec(fmt.Sprintf("delete from %s", sch.Table)).Error; err != nil {
			return fmt.Errorf("table:%s delete fail:%s", sch.Table, err.Error())
		}
	}
	return nil
}

func (x *XGorm) ClearAllData() {
	tmpDb := x.DB
	if tmpDb == nil {
		panic("尚未初始化数据库, 清空数据库失败")
	}
	tables := x.showTables()
	for _, table := range tables {
		if err := tmpDb.Exec(fmt.Sprintf("delete from %s", table)).Error; err != nil {
			panic("清空表数据失败:" + err.Error())
		}
	}
}

func (x *XGorm) showTables() []string {
	var sql string
	var tables []string
	switch x.config.DBType {
	case DBTypeMySQL:
		sql = "show tables"
	}
	if x.DB.Raw(sql).Scan(&tables).Error != nil {
		panic("表名获取失败")
	}
	return tables
}

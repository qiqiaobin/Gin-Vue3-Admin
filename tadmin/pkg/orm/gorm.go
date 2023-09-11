package orm

import (
	"fmt"
	"log"
	"strings"
	"time"

	"tadmin/conf"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// DB 数据库服务
var DB *gorm.DB

const (
	// maxOpenConns >= maxIdleConns

	defaultMaxOpenConns    = 100
	defaultMaxIdleConns    = 25
	defaultConnMaxLifetime = 10 * time.Minute
)

// Gorm 初始化数据库并产生数据库全局变量
func InitGorm() (*gorm.DB, error) {

	connStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Config.Database.User,
		conf.Config.Database.Password,
		conf.Config.Database.Host,
		conf.Config.Database.Port,
		conf.Config.Database.DbName,
	)

	var dialector gorm.Dialector

	switch strings.ToLower(conf.Config.Database.DBType) {
	case "mysql":
		dialector = mysql.Open(connStr)
	case "postgres":
		dialector = postgres.Open(connStr)
	default:
		return nil, fmt.Errorf("dialector(%s) not supported", conf.Config.Database.DBType)
	}

	gconfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.Config.Database.TablePrefix,
			SingularTable: true,
		},
	}
	db, err := gorm.Open(dialector, gconfig)
	if err != nil {
		return nil, err
	}

	if conf.Config.Database.Debug {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 验证与数据库的连接是否仍然有效，必要时建立连接。
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("[db ping failed] key name: %s, %w", conf.Config.Database.DbName, err)
	}

	maxOpenConns := defaultMaxOpenConns
	if conf.Config.Database.MaxOpenConns > 0 {
		maxOpenConns = conf.Config.Database.MaxOpenConns
	}

	maxIdleConns := defaultMaxIdleConns
	if conf.Config.Database.MaxIdleConns > 0 {
		maxIdleConns = conf.Config.Database.MaxIdleConns
	}

	if maxOpenConns < maxIdleConns {
		log.Fatalf("错误的数据库配置:%s, maxOpenConns需要大于等于 maxIdleConns, "+
			"将使用默认配置[defaultMaxOpenConns=%d, defaultMaxIdleConns=%d]",
			conf.Config.Database.DbName, defaultMaxOpenConns, defaultMaxIdleConns)
		maxOpenConns = defaultMaxOpenConns
		maxIdleConns = defaultMaxIdleConns
	}

	connMaxLifetime := defaultConnMaxLifetime
	if conf.Config.Database.MaxLifetime > 0 {
		if conf.Config.Database.MaxLifetime >= 60 {
			connMaxLifetime = time.Duration(conf.Config.Database.MaxLifetime) * time.Second
		} else {
			log.Fatalf("error config for database %s, connMaxLifetimeSeconds should be greater than 60 seconds"+
				"use the default [defaultConnMaxLifetime=%s]",
				conf.Config.Database.DbName, defaultConnMaxLifetime)
		}
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	// 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	sqlDB.SetMaxIdleConns(maxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	//设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDB.SetMaxOpenConns(maxOpenConns)

	// SetConnMaxLifetime 设置最大连接超时
	sqlDB.SetConnMaxLifetime(connMaxLifetime)

	//defer sqlDB.Close()

	fmt.Println("初始化数据库连接完成")
	// 赋值给私有全局变量
	//DB = db
	return db, nil
}

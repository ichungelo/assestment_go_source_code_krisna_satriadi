package gormconnection

import (
	"fmt"
	"log"
	"sync"

	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/config"
	appconfig "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/config/app_config"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	utillogger "github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils/util_logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB variable
var (
	DB   *gorm.DB
	lock *sync.Mutex = &sync.Mutex{}
)

// Create DB connection and migration using gorm
func connectDB(cfg *config.Config) *gorm.DB {
	dbName := fmt.Sprint(cfg.AppConfig.Name, "_", cfg.AppConfig.Stage)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", cfg.DBConfig.DbUsername, cfg.DBConfig.DbPassword, cfg.DBConfig.DbHost, cfg.DBConfig.DbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		utillogger.Error(err, nil)
		log.Fatal("Failed to run DB.\n", err.Error())
	}
	db.Begin()
	utillogger.Info("DB Connected", nil)

	if cfg.AppConfig.Stage == appconfig.DEV {
		db.Logger = logger.Default.LogMode(logger.Info)
	}

	createDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", dbName)
	db.Exec(createDatabaseCommand)

	utillogger.Info("running migrations...", nil)
	db.AutoMigrate(
		&model.Customer{},
		&model.Item{},
		&model.ItemType{},
		&model.Invoice{},
		&model.Quantity{},
	)
	utillogger.Info("migration finished!", nil)

	return db
}

func GetInstanceDB(cfg *config.Config) *gorm.DB {
	if DB == nil {
		lock.Lock()
		defer lock.Unlock()
		if DB == nil {
			utillogger.Info("Opering DB connection now.", nil)
			DB = connectDB(cfg)
		} else {
			utillogger.Info("DB connection already created.", nil)
		}
	} else {
		utillogger.Info("DB connection already created.", nil)
	}

	return DB
}

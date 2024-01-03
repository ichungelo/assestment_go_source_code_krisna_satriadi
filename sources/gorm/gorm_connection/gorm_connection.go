package gormconnection

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/config"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/utils"
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
func connectDB(stage config.Stage) *gorm.DB {
	var (
		host         = os.Getenv("DB_HOST")
		port         = os.Getenv("DB_PORT")
		user         = os.Getenv("DB_USER")
		pass         = os.Getenv("DB_PASS")
		dbNamePrefix = os.Getenv("APP_NAME")
		dbName       = fmt.Sprint(dbNamePrefix, "_", stage)
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", user, pass, host, port, dbName)
	
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		utils.Error(err, nil)
		log.Fatal("Failed to run DB.\n", err.Error())
	}
	db.Begin()
	utils.Info("DB Connected", nil)

	if stage == config.DEV {
		db.Logger = logger.Default.LogMode(logger.Info)
	}

	utils.Info("running migrations...", nil)
	db.AutoMigrate(
		&model.Customer{},
		&model.Item{},
		&model.ItemType{},
		&model.Invoice{},
		&model.Quantity{},
	)
	utils.Info("migration finished!", nil)

	return db
}

func GetInstanceDB(stage config.Stage) *gorm.DB {
	if DB == nil {
		lock.Lock()
		defer lock.Unlock()
		if DB == nil {
			utils.Info("Opering DB connection now.", nil)
			DB = connectDB(stage)
		} else {
			utils.Info("DB connection already created.", nil)
		}
	} else {
		utils.Info("DB connection already created.", nil)
	}

	return DB
}

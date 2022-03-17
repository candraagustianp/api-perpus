package database

import (
	"api-perpus/config"
	"errors"
	"fmt"

	log "github.com/siruspen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(conf *config.Config) *gorm.DB {
	if conf.TmpDep == "local" {
		db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.DBUsername, conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBName)))
		if err != nil {
			log.Fatalf("unable to initiate database connection: %v", err)
		}
		log.Infoln("Success initiate database connection")
		return db
	} else {
		db, err := gorm.Open(postgres.Open(config.GetString("DATABASE_URL")))
		if err != nil {
			log.Fatalf("unable to initiate database connection")
		}
		log.Infoln("Success to initiate database connection")
		return db
	}

}

func AutoMigrate(db *gorm.DB, name string, tabel interface{}) {
	if err := db.AutoMigrate(tabel); err != nil {
		log.Fatalf("cannot migrate a tabel: %v", err)
	} else {
		log.Printf("success migrate a table: %s", name)
	}
}

func GetAll(db *gorm.DB, value interface{}) error {
	if result := db.Order("id asc").Find(value); result.Error != nil || result.RowsAffected < 1 {
		if result.RowsAffected < 1 {
			return errors.New("there is no data")
		}
		return result.Error
	} else {
		return nil
	}
}

func GetWhere(db *gorm.DB, value interface{}, condition string) error {
	if result := db.Where(condition).Order("id asc").Find(value); result.Error != nil || result.RowsAffected < 1 {
		if result.RowsAffected < 1 {
			return errors.New("there is no data")
		}
		return errors.New("failed get data")
	} else {
		return nil
	}
}

func SaveData(db *gorm.DB, value interface{}) error {
	if result := db.Create(value); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func UpdateData(db *gorm.DB, condition string, value interface{}) error {
	if result := db.Where(condition).Updates(value); result.Error != nil || result.RowsAffected < 1 {
		return errors.New("failed update data: " + result.Error.Error())
	} else {
		return nil
	}
}

func DeleteData(db *gorm.DB, value interface{}, primary interface{}) error {
	if result := db.Delete(value, primary); result.Error != nil || result.RowsAffected < 1 {
		return errors.New("failed delete data: " + result.Error.Error())
	} else {
		return nil
	}
}

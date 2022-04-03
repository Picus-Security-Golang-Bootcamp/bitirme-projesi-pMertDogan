package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func ConnectPostgresDB() (*gorm.DB, error) {
	//get database connection string from env
	dataSourceName := fmt.Sprintf("host=%s port=%s  dbname=%s user=%s sslmode=disable password=%s",
	viper.GetString("database.host"),
	viper.GetString("database.port"),
	viper.GetString("database.dbname"),
	viper.GetString("database.username"),
	viper.GetString("database.password"),
	)
	
	// fmt.Println(viper.GetString("database.host"))
	// fmt.Println(viper.GetString("database.port"))
	// fmt.Println(viper.GetString("database.username"))
	// fmt.Println(viper.GetString("database.password"))

	//connect to DB
	gormDB2, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("cannot open database: %v", err)
	}
	//Maybe there is another way to get same reference to gormDB instance
	gormDB = gormDB2

	//get DB from gorm.DB
	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	//Check is connectio available
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return gormDB, nil
}


//Status Checker for DB
func StatusCheck() error {
	sqlDB, err := gormDB.DB()
	if err != nil {
		return err
	}
	//Check is connectio available
	if err := sqlDB.Ping(); err != nil {
		return err
	}

	fmt.Println(sqlDB.Stats())

	return nil
}

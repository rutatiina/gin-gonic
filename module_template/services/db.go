package services

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"wese/core/module_template/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnectAndMigrate() {

	// Open our jsonFile
	confFile, _err := os.Open("../conf.json")

	if _err != nil {
		panic("Failed to read conf file!")
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer confFile.Close()

	byteValue, _ := io.ReadAll(confFile)

	var config map[string]interface{}
	json.Unmarshal([]byte(byteValue), &config)

	// var dsn = "root:rutatiina@tcp(127.0.0.1:3306)/rg_gin_tonic?charset=utf8mb4&parseTime=true&loc=Local"
	var dsn = fmt.Sprint(config["mysql_username"]) + ":" + fmt.Sprint(config["mysql_password"]) + "@tcp(" + fmt.Sprint(config["mysql_db_path"]) + ")/" + fmt.Sprint(config["mysql_db_name"]) + "?charset=utf8mb4&parseTime=true&loc=Local"

	var db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&models.ModelName{})

	DB = db

}

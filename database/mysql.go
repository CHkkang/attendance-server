package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)
type DBConfig struct {
	Host  		string 	`yaml:"host"`
	Port     	int 	`yaml:"port"`
	User		string	`yaml:"user"`
	Password 	string	`yaml:"password"`
	DB       	string	`yaml:"database"`
}

type MysqlClient struct {
	DB  *gorm.DB
}

func ConnectDB(config *DBConfig) *MysqlClient {
	mysqlDatabase := new(MysqlClient)
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DB)

	mysqlDatabase.DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version}
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	//mysqlDatabase.DB.AutoMigrate(&models.User{}, &models.Image{})

	return mysqlDatabase
}

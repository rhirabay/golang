package infra

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Datasource struct {
	Connection *gorm.DB
}

func (datasource Datasource) GetConnection() *gorm.DB {
	if datasource.Connection != nil {
		return datasource.Connection
	}

	USER := "gorm"
	// PASS := os.Getenv("MYSQL_PASSWORD")
	PASS := "gorm"
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "gorm"

	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	connection, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connect to new database")

	return connection
}

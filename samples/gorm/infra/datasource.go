package infra

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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
	connection, err := gorm.Open("mysql", DSN)

	if err != nil {
		panic(err.Error())
	}

	return connection
}

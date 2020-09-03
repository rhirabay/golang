package main

import (
	"fmt"

	"github.com/rhirabay/golang/samples/gorm/domain"
	"github.com/rhirabay/golang/samples/gorm/infra"
)

func main() {
	fmt.Println("Hello world.")
	datasource := infra.Datasource{}

	connection := datasource.GetConnection()

	user := domain.User{ID: 1, Name: "User-1"}

	connection.Create(&user)
	fmt.Println(user)
}

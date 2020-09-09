package usecase

import (
	"fmt"

	"github.com/rhirabay/golang/samples/gorm/domain"
	"github.com/rhirabay/golang/samples/gorm/infra"
)

var datasource = infra.Datasource{}
var connection = datasource.GetConnection()

func CreateOne() {

	user := domain.User{Name: "User-1"}

	fmt.Println(user)
	connection.Create(&user)
	fmt.Println(user)
}

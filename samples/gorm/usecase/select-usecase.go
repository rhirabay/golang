package usecase

import (
	"fmt"

	"github.com/rhirabay/golang/samples/gorm/domain"
)

func FindSample() {
	users := []domain.User{}
	user := domain.User{Name: "User-1"}
	connection.Where(&user).Find(&users)
	// select id, name from users where name = 'User-1'
	fmt.Println(users)
	// [{4 User-1}]
}

func FindSampleStruct() {
	users := []domain.User{}
	user := domain.User{Name: "User-1"}
	connection.Where(&user).Find(&users)
	// select id, name from users where name = 'User-1'
	fmt.Println(users)
	// [{4 User-1}]
}

func FindSampleSQL() {
	users := []domain.User{}

	connection.Raw("select id, name from users where id = ?", 1).Scan(&users)
	// select id, name from users where id = 1
	fmt.Println(users)
	// [{1 User-4}]
}

func FindSampleWhere() {
	user := domain.User{}
	users := []domain.User{}

	connection.Where("id = ?", 1).First(&user)
	// select * from users where id = 1
	fmt.Println(user)
	// {1 User-4}

	connection.Where("id IN (?)", []int{1, 2}).Find(&users)
	// select * from users where id in (1, 2)
	fmt.Println(users)
	// [{1 User-4} {2 User-3}]

	connection.Where("id = ?", 1).Or("name = ?", "User-1").Find(&users)
	// 同じ：connection.Where("id = ? or name = ?", 1, "User-1").Find(&users)
	// select * from users where id = 1 or name = 'User-1'
	fmt.Println(users)
	// [{1 User-4} {4 User-1}]
}

func FindAllSample() {
	users := []domain.User{}
	connection.Find(&users)
	// select * from users
	fmt.Println(users)
	// [{1 User-4} {2 User-3} {3 User-2} {4 User-1}]
}

func FindByIdSample() {
	user := domain.User{}
	connection.First(&user, 1)
	// select * from users where id = 1
	fmt.Println(user)
	// {1 User-4}

	users := []domain.User{}
	connection.Find(&users, []int{1, 2})
	// select * from users where id in (1, 2)
	fmt.Println(users)
	// [{1 User-4} {2 User-3}]
}

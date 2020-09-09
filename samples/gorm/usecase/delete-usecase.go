package usecase

import (
	"errors"
	"fmt"

	"github.com/rhirabay/golang/samples/gorm/domain"
	"gorm.io/gorm"
)

func DeleteSample() {
	users := []domain.User{}
	connection.Find(&users)
	fmt.Println(users)

	connection.Transaction(func(tx *gorm.DB) error {
		// deleteOne(tx)
		// deleteWhere(tx)
		deleteSoft(tx)

		return errors.New("rollback")
	})
}

func deleteWhere(connection *gorm.DB) {
	// 削除
	connection.Where("id >= ?", 3).Delete(&domain.User{})
	// delete from users where id >= 3

	users := []domain.User{}
	connection.Find(&users)
	fmt.Println(users)
	// [{1 User-4} {2 User-3}]
}

func deleteOne(connection *gorm.DB) {
	// 削除
	user := domain.User{Id: 1}
	connection.Delete(&user)
	// delete from users where id = 1

	users := []domain.User{}
	connection.Find(&users)
	fmt.Println(users)
	// [{2 User-3} {3 User-2} {4 User-1}]
}

func deleteSoft(connection *gorm.DB) {
	user := domain.User{}
	connection.Where("id = ?", 1).First(&user)

	// 論理削除
	connection.Delete(&user)
	// update users set deleted_at="<now>" where id = 1

	connection.Unscoped().Where("id = ?", 1).First(&user)
	fmt.Println(user)
	// {1 User-4 {2020-09-08 22:21:43 +0900 JST true}}

	// 物理削除
	connection.Unscoped().Delete(&user)
	// delete from users where id = 1

	result := connection.Unscoped().Where("id = ?", 1).First(&user)
	fmt.Println(result.Error)
	// record not found
}

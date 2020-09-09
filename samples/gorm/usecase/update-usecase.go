package usecase

import (
	"fmt"

	"github.com/rhirabay/golang/samples/gorm/domain"
	"gorm.io/gorm"
)

func UpdateSample() {
	user := domain.User{}
	// gorm.Exprで計算式を記述
	connection.Model(&user).Update("name", gorm.Expr("concat('User-', id)"))
	// 全件取得
	users := []domain.User{}
	connection.Find(&users)
	fmt.Println(users)
	// [{1 User-1} {2 User-2} {3 User-3} {4 User-4}]
}

func UpdateSampleSQL() {
	// SQLを指定して更新
	connection.Exec("update users set name = 'User-X'")
	// 全件取得
	users := []domain.User{}
	connection.Find(&users)
	fmt.Println(users)
	// [{1 User-X} {2 User-X} {3 User-X} {4 User-X}]
}

func UpdateSampleBatch() {
	connection.Table("users").Where("name like ?", "User-%").Updates(domain.User{Name: "User-X"})

	users := []domain.User{}
	connection.Find(&users)
	fmt.Println(users)
}

func UpdateSampleMultiColumn() {
	user := domain.User{Id: 1}
	fmt.Println(user)

	// nameを「User-X」に更新
	result := connection.Model(&user).Update("name", "User-X")
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)

	connection.Model(&user).Update("name", "User-4")
}

func UpdateSampleMulti() {
	// 検索用にIDを指定
	userForSearch := domain.User{Id: 1}
	fmt.Println(userForSearch)

	// 更新したい内容を別の構造体に定義
	userForUpdate := domain.User{Name: "User-X"}
	// 同じ：map[string]interface{}{"name": "User-X"}

	// 更新
	connection.Model(&userForSearch).Updates(&userForUpdate)
	// update users set id = 1 where name = 'User-X'
	fmt.Println(userForSearch)

	userForUpdate = domain.User{Name: "User-4"}
	connection.Model(&userForSearch).Updates(&userForUpdate)
	fmt.Println(userForSearch)
}

func UpdateSampleUpdateOne() {
	user := domain.User{Id: 1}
	fmt.Println(user)

	// nameを「User-X」に更新
	connection.Model(&user).Update("name", "User-X")
	fmt.Println(user)
	// {1 User-X}

	connection.Model(&user).Update("name", "User-4")
	connection.Where("id = ?", 1).First(&user)
	fmt.Println(user)
}

func UpdateSampleSave() {
	user := domain.User{}
	// 「ID = 1」で検索
	connection.Where("id = ?", 1).First(&user)
	fmt.Println(user)
	// {1 User-4}

	// Nameを「User-X」に更新
	user.Name = "User-X"
	connection.Save(&user)

	// 「ID = 1」で検索
	connection.Where("id = ?", 1).First(&user)
	fmt.Println(user)
	// {1 User-X}

	connection.Model(&user).Update("name", "User-4")

	connection.Where("id = ?", 1).First(&user)
	fmt.Println(user)
}

package domain

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	DeletedAt gorm.DeletedAt
}

func (user User) AfterCreate() {
	fmt.Print("inserted User: ")
	fmt.Println(user)
}

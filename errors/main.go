package main

import (
	"database/sql"
	"errors"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	user User
)

//func Test() {
//
//	dsn := "root:123456@tcp(localhost:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
//	db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		fmt.Println("mysql connect error:",err)
//		return
//	}
//	//err := db.Raw("select * from users where id = ?", 2).Find(&user).Error
//	err = db.Where("id = ?", 2).First(&user).Error
//	fmt.Println(user)
//	fmt.Println(err)
//}

func main() {

	data, err := WrapQueryErr(1)
	err = errors.Unwrap(err)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println(sql.ErrNoRows)
		return
	}
	fmt.Println(data)

}

func WrapQueryErr(id int) (*User, error) {
	dsn := "root:123456@tcp(localhost:3306)/gin"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("mysql connect error:", err)
		return nil, err
	}
	defer db.Close()
	err = db.QueryRow("select * from users where id = ?", id).Scan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%w, data is nil", err)
		}
	}
	return &user, nil
}

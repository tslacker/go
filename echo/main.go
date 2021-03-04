package main

import (
	"fmt"
	"time"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int64  `xorm:"'id'"`
	Name     string `xorm:"'name'"`
	Password string `xorm:"'password'"`
}

type Person struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password`
}

func main() {
	start := time.Now()
	// engine, _ := xorm.NewEngine("mysql", "root:test@tcp([127.0.0.1]:3306)/test?charset=utf8mb4&parseTime=true")
	engine, _ := xorm.NewEngine("mysql", "root:test@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=true")

	_ = engine.Sync2(new(User))

	e := echo.New()
	// for j := 0; j < 10; j++ {
	// 	var hoge User
	// 	var user []User
	// 	for i := 0; i < 50000; i++ {
	// 	hoge.Name = "name"
	// 	hoge.Password = "huga"
	// 	user = append(user, hoge)
	// 	}
	// 	engine.Insert(&user)
	// }
	// var user []User
	// _ = engine.Table("user").Find(&user)
	// fmt.Println(user)
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
	})
	e.GET("/aaaa", func(c echo.Context) error {
		var user []User
		var hoge User
		var b Person
		_ = engine.Table("user").Find(&user)
		hoge.Name = "name"
		hoge.Password = "huga"
		b.Name = hoge.Name
		b.Password = hoge.Password
		b.Name = "name"
		b.Password = "huga"
		fmt.Println(user)
		return c.JSON(http.StatusOK, user)
	})
	e.Logger.Fatal(e.Start(":1313"))
}

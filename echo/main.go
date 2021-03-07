package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int64  `xorm:"pk autoincr not null" json:"id"`
	GameName string `xorm:"not null" json:"gamename"`
	Name     string `xorm:"not null" json:"name"`
	Password string `xorm:"not null" json:"password`
}

func main() {
	// engine, _ := xorm.NewEngine("mysql", "root:test@tcp([127.0.0.1]:3306)/test?charset=utf8mb4&parseTime=true")
	engine, _ := xorm.NewEngine("mysql", "root:test@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=true")

	_ = engine.Sync2(new(User))

	e := echo.New()

	e.POST("/insert", func(c echo.Context) error {
		post := new(User)
		c.Bind(post)
		engine.Insert(post)
		return c.JSON(http.StatusOK, post)
	})

	e.GET("/findall", func(c echo.Context) error {
		var user []User
		engine.Table("user").Find(&user)
		return c.JSON(http.StatusOK, user)
	})

	e.GET("/find", func(c echo.Context) error {
		u := new(User)
		err := c.Bind(u)
		if err != nil {
			return err
		}
		var user []User
		fmt.Println(u.GameName)
		engine.Table("user").Where("game_name = ?", u.GameName).Find(&user)
		return c.JSON(http.StatusOK, user)
	})

	e.POST("/update", func(c echo.Context) error {
		u := new(User)
		err := c.Bind(u)
		if err != nil {
			return err
		}
		engine.Table("user").Where("game_name = ?", u.GameName).Where("name = ?", u.Name).Update(u)
		return c.JSON(http.StatusOK, u)
	})

	e.Logger.Fatal(e.Start(":1313"))
}

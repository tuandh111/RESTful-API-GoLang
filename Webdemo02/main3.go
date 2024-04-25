package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

type TodoItem1 struct {
	Id          int        `json:"id" gorm:"id"`
	Title       string     `json:"title" gorm:"title"`
	Description string     `json:"description" gorm:"description"`
	Status      string     `json:"status" gorm:"status"`
	Create_at   *time.Time `json:"create_at" gorm:"create_at"`
	Update_at   *time.Time `json:"update_at" gorm:"update_at"`
}
type TodoitemInsert1 struct {
	Title       string `json:"title" gorm:"title"`
	Description string `json:"description" gorm:"description"`
	Status      string `json:"status" gorm:"status"`
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/todoitem?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Database is", db)
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		item := v1.Group("/item")
		{
			item.POST("", InsertTodoItem(db))
			item.GET("/:id", getItemById2(db))
		}
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello word",
		})
	})
	r.Run(":3000")

}
func InsertTodoItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data TodoitemInsert1
		if err := c.ShouldBind(&data); err != nil {
			log.Fatal(err.Error())
		}
		if err := db.Table("products").Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
func getItemById2(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data TodoItem1
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}
		db.Table("products").Where("id =?", id).First(&data)
	}
}

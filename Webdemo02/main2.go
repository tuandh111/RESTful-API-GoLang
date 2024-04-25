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

type TodoItem2 struct {
	Id          int        `json:"id" gorm:"id"`
	Title       string     `json:"title" gorm:"title"`
	Description string     `json:"description" gorm:"description"`
	Status      string     `json:"status" gorm:"status"`
	Create_at   *time.Time `json:"create_at" gorm:"create_at"`
	Update_at   *time.Time `json:"update_at" gorm:"update_at"`
}
type TodoItem2Insert struct {
	Title       string `json:"title" gorm:"title"`
	Description string `json:"description" gorm:"description"`
	Status      string `json:"status" gorm:"status"`
}

func main() {
	fmt.Println("Hello word")
	dsn := "root:123456@tcp(127.0.0.1:3306)/todoitem?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal()
	}
	fmt.Println("Database is", db)
	r := gin.Default()
	item := r.Group("v1")
	{
		v1 := item.Group("item")
		{
			v1.POST("", InsertItem(db))
			v1.GET("/:id", GetItemByidMain2(db))
			v1.GET("/list", GetAllListMain2(db))
		}
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello word",
		})
	})
	r.Run(":3000")
}
func InsertItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data TodoItem2Insert
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}
		if err := db.Table("products").Create(&data).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"err": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": &data,
		})
	}
}
func GetItemByidMain2(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data TodoItem2
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"err": err.Error(),
			})
			return
		}
		if err := db.Table("products").Where("id = ?", id).First(&data).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"Err": err.Error,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
func GetAllListMain2(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data []TodoItem2
		if err := db.Table("products").Find(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

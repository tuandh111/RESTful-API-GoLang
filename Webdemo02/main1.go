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

type TodoItem struct {
	Id          int        `json:"id" gorm:"id"`
	Title       string     `json:"title" gorm:"title"`
	Description string     `json:"description" gorm:"description"`
	Status      string     `json:"status" gorm:"status"`
	Update_at   *time.Time `json:"update_at" gorm:"update_at"`
	Create_at   *time.Time `json:"create_at" gorm:"create_at"`
}
type TodoitemInsert struct {
	Title       string `json:"title" gorm:"title"`
	Description string `json:"description" gorm:"description"`
	Status      string `json:"status" gorm:"status"`
}
type TodoitemUpdate struct {
	Title       string  `json:"title" gorm:"title"`
	Description string  `json:"description" gorm:"description"`
	Status      *string `json:"status" gorm:"status"`
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/todoitem?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	db.AutoMigrate(&TodoItem{})
	fmt.Println("Database is: ", db)

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		items := v1.Group("/item")
		{
			items.GET("/list", GetAllList(db))
			items.POST("", CreateDb(db))
			items.GET("/:id", GetItemById1(db))
			items.DELETE("/:id", DeleteById1(db))
			items.PATCH("/:id", UpdateById1(db))
		}
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000")
}
func CreateDb(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data TodoitemInsert
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}
		if err := db.Table("todo_items").Create(&data).Error; err != nil {
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

func GetItemById1(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data TodoItem
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": "Đây không phải là kiểu id",
			})
			return
		}
		if err := db.Table("todo_items").Where("id = ?", id).First(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Data": data,
		})
	}
}
func UpdateById1(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data TodoitemUpdate
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"err": err.Error(),
			})
			return
		}
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}
		if err := db.Table("todo_items").Where("id= ?", id).Updates(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"Trạng thái": data,
		})
	}
}
func DeleteById1(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data TodoItem
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}
		if err := db.Table("todo_items").Where("id = ?", id).Delete(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Delete is succefully",
		})
	}
}
func GetAllList(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data []TodoItem
		if err := db.Table("todo_items").Find(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"err": data,
		})
	}
}

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

type Products struct {
	Id          int        `json:"id" gorm:"id"`
	Title       string     `json:"title" gorm:"title"`
	Description string     `json:"description" gorm:"title"`
	Status      string     `json:"status" gorm:"status"`
	Create_at   *time.Time `json:"create_at" gorm:"create_at"`
	Update_at   *time.Time `json:"update_at" gorm:"update_at"`
}
type TodoItemcreated struct {
	Title       string `json:"title" gorm:"title"`
	Description string `json:"description" gorm:"description"`
	Status      string `json:"status" gorm:"status"`
}

func (Products) tableName() string {
	return "product"
}
func main() {
	fmt.Println("hello word")
	dsn := "root:123456@tcp(127.0.0.1:3306)/todoitem?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Database is:", db)
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		item := v1.Group("/item")
		{
			item.POST("", Insertdb(db))
			item.GET("/:id", GetItemById(db))
			item.GET("/list", GetAllItems(db))
			item.PATCH("/:id", UpdateById(db))
			item.DELETE("/:id", DeleteById(db))

		}
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello word",
		})
	})
	r.Run(":3000")

}
func Insertdb(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data TodoItemcreated
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error1": err.Error(),
			})
		}
		if err := db.Table(Products{}.tableName()).Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error1": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})

	}
}
func GetItemById(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data Products
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Id không phải là số",
			})
			return
		}
		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": data,
		})

	}
}
func UpdateById(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data Products
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Id không phải là số",
			})
			return
		}
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error1": err.Error(),
			})
		}
		if err := db.Table("products").Where("id = ?", id).Updates(&data).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Update succesfully",
		})
	}
}

func DeleteById(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data Products
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Id không phải là số",
			})
			return
		}
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error1": err.Error(),
			})
		}
		if err := db.Table("products").Where("id = ?", id).Delete(&data).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Deleted succesfully",
		})
	}
}

func GetAllItems(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var products []Products

		if err := db.Find(&products).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": products,
		})
	}
}

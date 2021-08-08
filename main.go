package main

import (
	"context"
	"database/sql"
	"log"
	"github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	"net/http"
	"gorm.io/gorm"
)

var db *sql.DB

var server = "localhost"
var port = 1433
var user = "sa"
var password = "123456789"
var database = "restaurant"

func main() {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	var err error

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")
	if err:=runService(db) ;err!=nil {
		log.Fatal(err)
	}

}
func runService(db *sql.DB) error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

restaurants :=r.Group("/restaurants")
{
	restaurants.Post("", func(c *gin.Context) {
		var data Restaurant
		if err:=c.ShouldBind(&data); err!=nil {
			c.JSON(http.StatusOK, map[string]interface{}{
				"error":err.Error(),
			})
			return
		}
		if err:=db.Create(&data).Error;err !=nil{
			c.JSON(http.StatusOK, map[string]interface{}{
				"error":err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, data)
	})
	restaurants.GET("/id",func(c *gin.Context){
		id,err :=strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusOK,map[string]interface{}{
				"error":err.Error(),
			})
			return
		}
		var data Restaurant
		if err :=db.Where("id =?",id).First(&data).Error;err !=nil{
			c.JSON(http.StatusOK,map[string]interface{}{
				"error":err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,data)
	})
	restaurants.GET("",func(c *gin.Context){
		var data []Restaurant
		type Filter struct{
			CityId int `json:"city_Id" form:"city_Id"`
		}
		var filter Filter
		c.ShouldBind(&filter)
		newDB :=db
		if filter.CityId >0{
			newDB =db.Where("city_Id =?" ,Filter.CityId)
		}
		if err :=newDB.Find(&data).Error;err !=nil{
			c.JSON(http.StatusOK, map[string]interface{}{
				"error" :err.Error(),
			})
			return 
		}
		c.JSON(http.StatusOK, data)
	})
}
type Restaurant struct {
	Id int `json:"id" gorm:"comlumn:id;"`
	Name string `json:"name" gorm:"comlumn:name;"`
	Addr string `json:"address" gorm:"comlumn:addr;"`
}
func (Restaurant) TableName() string {
	return "restaurants"
}

	



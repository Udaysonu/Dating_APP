package server

import (
	"github.com/udaysonu/Dating-Backend/internal/dao/models"
	ginSwagger "github.com/swaggo/gin-swagger"   
	"github.com/swaggo/gin-swagger/swaggerFiles"	
	_ "github.com/udaysonu/Dating-Backend/docs"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/udaysonu/Dating-Backend/database"
	"github.com/udaysonu/Dating-Backend/config"
	"github.com/gin-contrib/cors"
	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
 )

var db *gorm.DB

var err error

func LoadUsersData(){
	jsonFile, err := os.Open("../../data/users.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users []models.User
	
	json.Unmarshal(byteValue, &users)
	
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i].Name,users[i].Email)
		db.Create(&users[i])
	}
		
}
func LoadLikesData(){
	jsonFile, err := os.Open("../../data/likes.json")

	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var likes []models.Like
	
	json.Unmarshal(byteValue, &likes)
	
	for i := 0; i < len(likes); i++ {
		fmt.Println(likes[i].FromId,likes[i].ToId)
		db.Create(&likes[i])
	}
		
}


func RunServer()error{
 	server:=gin.Default()
    server.Use(cors.Default())

	db=database.GetInstance()
 
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Like{})

// uncomment these lines if u r running the code for the first time
// 	LoadUsersData()
// 	LoadLikesData()

	datingRoute:=server.Group("/dating")
	{
		DatingRoute(datingRoute.Group("/"))
	}

	server.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
 	err:=server.Run(":"+config.Server["PORT"])
	log.Info("Server starting on PORT: ",config.Server["PORT"])

	if err!=nil{
		log.WithField("Error: ", err).Error("error while starting the server")
		return err
	}

	return nil;
}

 
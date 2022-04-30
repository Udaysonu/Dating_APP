package database
import (
	"github.com/udaysonu/Dating-Backend/config"
	"github.com/jinzhu/gorm"
	"fmt"
)


var db *gorm.DB
var err error

func GetInstance()*gorm.DB{

	if db==nil{
		dbURI:=fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",config.Database["SITE"],config.Database["PORT"],config.Database["USER"],config.Database["DBNAME"],config.Database["PASSWORD"])
		db, err = gorm.Open( "postgres", dbURI)
	
		if err != nil {

			fmt.Println("ERROR:",err)
			panic("failed to connect database")

		}

 	}

	return db
}


func CloseInstance(){
	if db!=nil{
		db.Close()
	}
}
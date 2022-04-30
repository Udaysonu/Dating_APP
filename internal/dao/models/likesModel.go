package models

import (
	_ "github.com/lib/pq"
)


type Like struct{
	
	Id uint	  `gorm:"primaryKey;autoIncrement;not_null"  json:"id" `  
	FromId uint	`json:"who_likes"`
	ToId   uint	`json:"who_is_liked"`

}

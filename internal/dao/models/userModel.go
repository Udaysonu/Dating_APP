package models

import (
	_ "github.com/lib/pq"
)

// gorm.Model
type User struct{
	
	Id uint	  `gorm:"primaryKey;autoIncrement;not_null"  json:"id" `  
	Name string	  `json:"name"`
	Email string  `json:"email"`
	Location int  `json:"location"`
	Gender string `json:"gender"`
}

//`gorm:"typevarchar(100);unique_index" json:"email"`
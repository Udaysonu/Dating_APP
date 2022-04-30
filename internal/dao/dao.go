package dao
import (
	"github.com/udaysonu/Dating-Backend/internal/dao/models"
	"github.com/udaysonu/Dating-Backend/database"
	"github.com/jinzhu/gorm"
)
var user models.User
var users []models.User
var like models.Like
var likes []models.Like
var db *gorm.DB

type DatingDao struct{
	Db *gorm.DB
}
type IDao interface{
	FindMatch(id string)([]models.User,error)
	GetAllMatches()([]models.Pair,error)
	GetUsersWithinRange(id string,rang int)([]models.User,error)
	GetAllUsersWithAlphabets(match string)([]models.User,error)
}

func GetDatingDao() IDao{
	return &DatingDao{
		database.GetInstance(),
	}
}


var response = []models.Response{}

func (dd *DatingDao)GetAllMatches()([]models.Pair,error){
	resp:=[]models.Pair{}
	result:=dd.Db.Raw("SELECT u1.from_id,u1.to_id FROM likes u1 INNER JOIN likes u2 ON u1.from_id=u2.to_id AND u1.to_id=u2.from_id WHERE u1.from_id>u2.from_id").Scan(&response)
	if result.Error!=nil{
		return resp,result.Error
	}
	for _,val:=range response{
		var user1 models.User
		var user2 models.User
		dd.Db.Find(&user1,val.FromId)
		dd.Db.Find(&user2,val.ToId)
		resp=append(resp,models.Pair{user1,user2})
	}
	return resp,result.Error
}

func (dd *DatingDao)GetUsersWithinRange(id string,rang int)([]models.User,error){
	result:=dd.Db.Find(&user,id)
	if	result.Error!=nil{
		return users,nil;
	}
	result=dd.Db.Raw("SELECT * FROM users WHERE ABS(users.location-?)<=?",user.Location,rang).Scan(&users)
	return users,result.Error
}

func (dd *DatingDao)GetAllUsersWithAlphabets(match string)([]models.User,error){
	result:=dd.Db.Where("LOWER(name) LIKE ?", "%"+match+"%").Find(&users)
	return users,result.Error
}

func (dd *DatingDao)FindMatch(id string)([]models.User,error){
	result:=dd.Db.Raw("SELECT * FROM users WHERE users.id IN (SELECT likes.to_id FROM likes WHERE likes.from_id=? AND likes.to_id IN (SELECT likes.from_id FROM likes WHERE likes.to_id=?))",id,id).Scan(&users)
	return users,result.Error
}







// func GetUserById(){
// 	db=database.GetInstance()
// 	result:=db.Find(&user,1)
// 	fmt.Println(user,result.Error)
// }

// func GetAllLikers(){
// 	db=database.GetInstance()
// 	db.Where("to_id = ?",66).Find(&likes)
// 	fmt.Println(likes)
// }

// func FindMatch()[]models.{
// 	db=database.GetInstance()
// 	db.Where("to_id = ?",66).Find(&likes)
// 	answers:=[]models.Like{}
// 	var answer models.Like
// 	for _,val:=range likes{
// 		result:=db.Where("to_id = ? AND from_id = ?",val.FromId,66).Find(&answer)
// 		if result.Error==nil{
// 			answers=append(answers,answer)
// 		}
// 	}
// 	return answers
// }

// func main(){
// 	GetMatch()
// }
package services

import (
	"github.com/udaysonu/Dating-Backend/internal/dao/models"
	// "github.com/udaysonu/Dating-Backend/database"
	"github.com/udaysonu/Dating-Backend/internal/dao"
)


type IService interface{
	GetMatchService(id string)([]models.User,error)
	GetAllMatchesService()([]models.Pair,error)
	GetUsersWithAlphabetService(match string)([]models.User,error)
	GetUsersWithRangeService(id string,rang int)([]models.User,error)
}

type DatingService struct{
	Dao dao.IDao
}

func NewDatingService() IService{
	return &DatingService{
		dao.GetDatingDao(),
	}
}

func (ds *DatingService)GetMatchService(id string)([]models.User,error){
	return ds.Dao.FindMatch(id)
}

func (ds *DatingService)GetAllMatchesService()([]models.Pair,error){
	return ds.Dao.GetAllMatches()
}

func (ds *DatingService)GetUsersWithAlphabetService(match string)([]models.User,error){
	return ds.Dao.GetAllUsersWithAlphabets(match)
}

func (ds *DatingService)GetUsersWithRangeService(id string,rang int)([]models.User,error){
	return ds.Dao.GetUsersWithinRange(id,rang)
}
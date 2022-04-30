package handler

import (
	"strings"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/udaysonu/Dating-Backend/internal/dao/models"
	"github.com/udaysonu/Dating-Backend/internal/server/services"
)

type DatingController struct {
	Service services.IService
}

func NewDatingController() ICustomer {
	return &DatingController{
		services.NewDatingService(),
	}
}


type ICustomer interface {
	GetMatch(ctx *gin.Context) ([]models.User,error)
	GetAllMatches(ctx *gin.Context) ([]models.Pair,error)
	GetUsersWithAlphabet(ctx *gin.Context) ([]models.User,error)
	GerUsersWithRange(ctx *gin.Context)  ([]models.User,error)
}


func (dc *DatingController)GetMatch(ctx *gin.Context)([]models.User,error){
	id:=ctx.Param("id");
	return dc.Service.GetMatchService(id)
}

func (dc *DatingController)GetAllMatches(ctx *gin.Context)([]models.Pair,error){
	return dc.Service.GetAllMatchesService()	
}

func (dc *DatingController)GetUsersWithAlphabet(ctx *gin.Context)([]models.User,error){
	match:=ctx.Param("match")
	match=strings.ToLower(match)
	return dc.Service.GetUsersWithAlphabetService(match)
}

func (dc *DatingController)GerUsersWithRange(ctx *gin.Context)([]models.User,error){
	rang:=ctx.Param("range")
	id:=ctx.Param("id")
	rang1,err:=strconv.Atoi(rang)
	if err!=nil{
		return []models.User{},err
	}
	return dc.Service.GetUsersWithRangeService(id,rang1)
}




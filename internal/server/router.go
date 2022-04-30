package server
import(
	"github.com/gin-gonic/gin"
	"github.com/udaysonu/Dating-Backend/internal/server/handler"
)


func DatingRoute(g *gin.RouterGroup){
	hndler:=handler.NewDatingController()
	g.GET("/match/:id",func(ctx *gin.Context){
		result,err:=hndler.GetMatch(ctx)
		if err==nil{
			ctx.JSON(200,gin.H{"data":result,"message":"Successfully fetched"})
		} else {
			ctx.JSON(400,gin.H{"data":nil,"message":err.Error()})
		}
	})

	g.GET("/allmatches",func(ctx *gin.Context){
		result,err:=hndler.GetAllMatches(ctx)
		if err==nil{
			ctx.JSON(200,gin.H{"data":result,"message":"Successfully fetched"})
		} else {
			ctx.JSON(400,gin.H{"data":nil,"message":err.Error()})
		}
	})

	g.GET("/matchalpha/:match",func(ctx *gin.Context){
		result,err:=hndler.GetUsersWithAlphabet(ctx)
		if err==nil{
			ctx.JSON(200,gin.H{"data":result,"message":"Successfully fetched"})
		} else {
			ctx.JSON(400,gin.H{"data":nil,"message":err.Error()})
		}
	})

	g.GET("/range/:id/:range",func(ctx *gin.Context){
		result,err:=hndler.GerUsersWithRange(ctx)
		if err==nil{
			ctx.JSON(200,gin.H{"data":result,"message":"Successfully fetched"})
		} else {
			ctx.JSON(400,gin.H{"data":nil,"message":err.Error()})
		}
	})

}
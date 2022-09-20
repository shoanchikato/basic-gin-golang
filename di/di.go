package di

import (
	"basic-gin/db"
	"basic-gin/repo"
	"basic-gin/route"

	"github.com/gin-gonic/gin"
)

func NewApp() *gin.Engine {
	db := db.InitDB()
	r := gin.Default()

	// repo
	prp := repo.NewPeopleRepo(db)

	// routers
	p := r.Group("/people")

	// routes
	prt := route.NewPeopleRoutes(prp)

	// people routes
	p.GET("", prt.GetAll())
	p.GET("/:id", prt.GetOne())
	p.POST("", prt.Create())
	p.PUT("/:id", prt.Update())

	return r
}

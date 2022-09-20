package route

import (
	"basic-gin/model"
	"basic-gin/repo"

	"github.com/gin-gonic/gin"
	"net/http"
)

type PeopleRoute interface {
	Create() func(c *gin.Context)
	GetOne() func(c *gin.Context)
	GetAll() func(c *gin.Context)
	Update() func(c *gin.Context)
}

type peopleRoute struct {
	repo repo.PeopleRepo
}

func NewPeopleRoutes(repo repo.PeopleRepo) PeopleRoute {
	return &peopleRoute{repo}
}

// Create
func (p *peopleRoute) Create() func(c *gin.Context) {
	return func(c *gin.Context) {
		person := model.Person{}

		if err := c.ShouldBindJSON(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		p.repo.Create(&person)

		c.JSON(http.StatusCreated, person)
	}
}

// GetAll
func (p *peopleRoute) GetAll() func(c *gin.Context) {
	return func(c *gin.Context) {
		people := p.repo.GetAll()

		c.JSON(http.StatusOK, people)
	}
}

// GetOne
func (p *peopleRoute) GetOne() func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := getIDParam(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		person := p.repo.GetOne(uint(id))

		c.JSON(http.StatusOK, person)
	}
}

// Update
func (p *peopleRoute) Update() func(c *gin.Context) {
	return func(c *gin.Context) {
		person := model.Person{}

		if err := c.ShouldBindJSON(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := getIDParam(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		
		p.repo.Update(id, &person)

		c.JSON(http.StatusOK, person)
	}
}

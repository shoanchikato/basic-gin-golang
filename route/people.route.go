package route

import (
	"basic-gin/model"
	"basic-gin/repo"

	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

		c.JSON(201, person)
	}
}

// GetAll
func (p *peopleRoute) GetAll() func(c *gin.Context) {
	return func(c *gin.Context) {
		people := p.repo.GetAll()

		c.JSON(200, people)
	}
}

// GetOne
func (p *peopleRoute) GetOne() func(c *gin.Context) {
	return func(c *gin.Context) {
		idString := c.Param("id")

		id, err := strconv.Atoi(idString)
		if err != nil {
			c.String(400, "bad id value, can't be parsed to an interger")
		}

		person := p.repo.GetOne(uint(id))

		c.JSON(200, person)
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

		idString := c.Param("id")

		idInt, err := strconv.Atoi(idString)
		if err != nil {
			c.String(400, "bad id value, can't be parsed to an interger")
		}

		id := uint(idInt)
		p.repo.Update(id, &person)

		c.JSON(200, person)
	}
}

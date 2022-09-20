package repo

import (
	"basic-gin/model"

	"gorm.io/gorm"
)

type PeopleRepo interface {
	Create(person *model.Person) *model.Person
	Update(id uint, person *model.Person) *model.Person
	GetOne(id uint) *model.Person
	GetAll() *[]model.Person
}

type peopleRepo struct {
	db *gorm.DB
}

func NewPeopleRepo(db *gorm.DB) PeopleRepo {
	return &peopleRepo{db}
}

// Create
func (p *peopleRepo) Create(person *model.Person) *model.Person {
	p.db.Create(person)

	return person
}

// GetAll
func (p *peopleRepo) GetAll() *[]model.Person {
	people := new([]model.Person)
	p.db.Find(people)

	return people
}

// GetOne
func (p *peopleRepo) GetOne(id uint) *model.Person {
	person := new(model.Person)
	p.db.Find(&person, "id = ?", id)

	return person
}

// Update
func (p *peopleRepo) Update(id uint, person *model.Person) *model.Person {
	person.ID = id
	p.db.Save(person)

	return person
}

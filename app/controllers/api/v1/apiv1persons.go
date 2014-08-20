package controllers

import (
	"revel_example/app/models"

	"github.com/revel/revel"
)

type APIV1Persons struct {
	*revel.Controller
}

func (c APIV1Persons) Index() revel.Result {
	persons := models.AllPersons()

	return c.RenderJson(persons)
}

func (c APIV1Persons) Show(id int64) revel.Result {
	person := models.FindPerson(id)

	return c.RenderJson(person)
}

func (c APIV1Persons) Create(person models.Person) revel.Result {
	models.CreatePerson(person)

	response := make(map[string]string)
	response["message"] = "Record created successfuly."

	return c.RenderJson(response)
}

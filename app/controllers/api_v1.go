package controllers

import (
	"revel_example/app/models"

	"github.com/revel/revel"
)

type API_V1 struct {
	*revel.Controller
}

func (c API_V1) Index() revel.Result {
	persons := models.AllPersons()

	return c.RenderJson(persons)
}

func (c API_V1) Show(id int64) revel.Result {
	person := models.FindPerson(id)

	return c.RenderJson(person)
}

func (c API_V1) Create(person models.Person) revel.Result {
	models.CreatePerson(person)

	response := make(map[string]string)
	response["message"] = "Record created successfully."

	return c.RenderJson(response)
}

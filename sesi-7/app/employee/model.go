package employee

import "github.com/google/uuid"

type Employee struct {
	Id     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Salary int       `json:"salary"`
	Age    int       `json:"age"`
}

type SearchEmployee struct {
	Query string `json:"query"`
}

type FilterableAttributes struct {
	Fields []string `json:"fields"`
}

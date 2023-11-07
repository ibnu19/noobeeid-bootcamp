package employee

import "github.com/google/uuid"

type EmployeeRequest struct {
	Id     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Salary int       `json:"salary"`
	Age    int       `json:"age"`
}

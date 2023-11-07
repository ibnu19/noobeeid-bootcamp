package employee

import (
	"log"
	"sesi-7/pkg"

	"github.com/gofiber/fiber/v2"
)

func RegisterEmployeeService(router fiber.Router) {
	db, err := pkg.InitDB()
	if err != nil {
		log.Println(err.Error())
	}

	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	employeeRouter := router.Group("/employees")
	{
		employeeRouter.Post("", handler.CreateNewEmployee)
		employeeRouter.Put("", handler.UpdateEmployee)
		employeeRouter.Delete("/:id", handler.DeleteEmployee)
		employeeRouter.Get("/:id", handler.GetOne)
		employeeRouter.Post("/fetch", handler.GetAll)
		employeeRouter.Get("", handler.SearchQuery)
		employeeRouter.Post("/search", handler.Search)
		employeeRouter.Put("/settings/filterable-attributes", handler.UpdateFilterableAttributes)
	}
}

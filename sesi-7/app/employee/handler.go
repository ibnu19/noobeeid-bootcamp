package employee

import (
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service
}

func NewHandler(service service) handler {
	return handler{
		service: service,
	}
}

func (handler *handler) CreateNewEmployee(c *fiber.Ctx) error {
	request := EmployeeRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(ApiResponse(fiber.ErrBadRequest.Message, nil, err))
	}

	response, err := handler.service.Create(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(ApiResponse(fiber.ErrBadRequest.Message, response.Status, err))
	}
	return c.Status(fiber.StatusOK).JSON(ApiResponse("create success", nil, nil))
}

func (handler *handler) UpdateEmployee(c *fiber.Ctx) error {
	request := EmployeeRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(ApiResponse(fiber.ErrBadRequest.Message, nil, err))
	}

	response, err := handler.service.Update(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(ApiResponse(fiber.ErrBadRequest.Message, response.Status, err))
	}
	return c.Status(fiber.StatusOK).JSON(ApiResponse("update success", nil, nil))
}

func (handler *handler) DeleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	employee, err := handler.service.GetOne(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).
			JSON(ApiResponse(fiber.ErrNotFound.Message, nil, err))
	}

	response, err := handler.service.Delete(employee.(EmployeeRequest))
	if err != nil {
		return c.Status(fiber.StatusNotFound).
			JSON(ApiResponse(fiber.ErrNotFound.Message, response.Status, err))
	}
	return c.Status(fiber.StatusOK).JSON(ApiResponse("delete success", nil, nil))
}

func (handler *handler) GetOne(c *fiber.Ctx) error {
	id := c.Params("id")
	response, err := handler.service.GetOne(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).
			JSON(ApiResponse(fiber.ErrNotFound.Message, nil, err))
	}
	return c.Status(fiber.StatusOK).JSON(ApiResponse("success", response, nil))
}

func (handler *handler) GetAll(c *fiber.Ctx) error {
	responses, err := handler.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusNotFound).
			JSON(ApiResponse(fiber.ErrNotFound.Message, nil, err))
	}
	return c.Status(fiber.StatusOK).JSON(ApiResponse("success", responses.Results, nil))
}

func (handler *handler) SearchQuery(c *fiber.Ctx) error {
	req := c.Query("q")
	response, err := handler.service.Search(req)
	if err != nil {
		return c.Status(fiber.StatusNotFound).
			JSON(ApiResponse(fiber.ErrNotFound.Message, nil, err))
	}
	return c.Status(fiber.StatusOK).JSON(ApiResponse("success", response, nil))
}

func (handler *handler) Search(c *fiber.Ctx) error {
	var req SearchEmployee
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(ApiResponse(fiber.ErrBadRequest.Message, nil, err))
	}

	response, err := handler.service.Search(req.Query)
	if err != nil {
		return c.Status(fiber.StatusNotFound).
			JSON(ApiResponse(fiber.ErrNotFound.Message, nil, err))
	}
	return c.Status(fiber.StatusOK).JSON(ApiResponse("success", response, nil))
}

func (handler *handler) UpdateFilterableAttributes(c *fiber.Ctx) error {
	var request FilterableAttributes
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(ApiResponse(fiber.ErrBadRequest.Message, nil, err))
	}

	response, err := handler.service.UpdateFilterableAttributes(request.Fields)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(ApiResponse(fiber.ErrBadRequest.Message, nil, err))
	}
	return c.Status(fiber.StatusOK).JSON(ApiResponse("success", response, nil))
}

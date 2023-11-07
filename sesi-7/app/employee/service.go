package employee

import (
	"github.com/google/uuid"
	"github.com/meilisearch/meilisearch-go"
)

type service struct {
	repository
}

func NewService(repository repository) service {
	return service{
		repository: repository,
	}
}

func (service *service) Create(request EmployeeRequest) (response *meilisearch.TaskInfo, err error) {
	employee := Employee{}
	request.Id = uuid.New()
	RequestToStructConvert(request, &employee)
	employees := []Employee{
		employee,
	}
	return service.repository.Save(employees)
}

func (service *service) Update(request EmployeeRequest) (response *meilisearch.TaskInfo, err error) {
	employee := Employee{}
	RequestToStructConvert(request, &employee)
	employees := []Employee{
		employee,
	}
	return service.repository.Update(employees)
}

func (service *service) Delete(request EmployeeRequest) (response *meilisearch.TaskInfo, err error) {
	return service.repository.Delete(request.Id.String())
}

func (service *service) GetOne(id string) (response any, err error) {
	return service.repository.FindById(id)
}

func (service *service) GetAll() (response meilisearch.DocumentsResult, err error) {
	return service.repository.FindAll()
}

func (service *service) Search(req string) (employees []Employee, err error) {
	return service.repository.Search(req)
}

func (service *service) UpdateFilterableAttributes(fields []string) (response *meilisearch.TaskInfo, err error) {
	return service.repository.UpdateFilterableAttributes(fields)
}

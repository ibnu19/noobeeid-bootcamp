package employee

import (
	"encoding/json"

	"github.com/meilisearch/meilisearch-go"
)

type repository struct {
	client *meilisearch.Client
}

func NewRepository(client *meilisearch.Client) repository {
	return repository{
		client: client,
	}
}

func (repository *repository) Save(employees []Employee) (response *meilisearch.TaskInfo, err error) {
	documents := StructToMapConvert(employees)
	return repository.client.Index("employees").AddDocuments(documents, "id")
}

func (repository *repository) Update(employees []Employee) (response *meilisearch.TaskInfo, err error) {
	documents := StructToMapConvert(employees)
	return repository.client.Index("employees").UpdateDocuments(documents)
}

func (repository *repository) Delete(id string) (response *meilisearch.TaskInfo, err error) {
	return repository.client.Index("employees").DeleteDocument(id)
}

func (repository *repository) FindById(id string) (response Employee, err error) {
	repository.client.Index("employees").GetDocument(id, &meilisearch.DocumentQuery{}, &response)
	return
}

func (repository *repository) FindAll() (responses meilisearch.DocumentsResult, err error) {
	repository.client.Index("employees").GetDocuments(&meilisearch.DocumentsQuery{
		Limit: 50,
	}, &responses)
	return
}

func (repository *repository) Search(req string) (employees []Employee, err error) {
	response, err := repository.client.Index("employees").Search(req, &meilisearch.SearchRequest{
		Limit:  10,
		Filter: "age >= 30",
	})
	if err != nil {
		return
	}

	hitByte, err := json.Marshal(response.Hits)
	if err != nil {
		return
	}

	err = json.Unmarshal(hitByte, &employees)
	if err != nil {
		return
	}

	return
}

func (repository *repository) UpdateFilterableAttributes(fields []string) (response *meilisearch.TaskInfo, err error) {
	repository.client.Index("employees").UpdateFilterableAttributes(&fields)
	return
}

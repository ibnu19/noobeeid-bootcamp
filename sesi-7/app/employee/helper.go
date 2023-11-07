package employee

import (
	"encoding/json"
	"log"

	"github.com/meilisearch/meilisearch-go"
)

func StructToMapConvert(object any) (result []map[string]any) {
	objectByte, err := json.Marshal(object)
	if err != nil {
		log.Println(err.Error())
	}

	err = json.Unmarshal(objectByte, &result)
	if err != nil {
		log.Println(err.Error())
	}
	return
}

func ResponseToStruct(response *meilisearch.SearchResponse, object *[]Employee) {
	respByte, err := json.Marshal(response)
	if err != nil {
		log.Println(err.Error())
	}

	err = json.Unmarshal(respByte, &object)
	if err != nil {
		log.Println(err.Error())
	}
}

func RequestToStructConvert(req EmployeeRequest, payload *Employee) {
	payload.Id = req.Id
	payload.Name = req.Name
	payload.Salary = req.Salary
	payload.Age = req.Age
}

package controller

import (
	"encoding/json"
	"github.com/bdxygy/exception"
	"github.com/bdxygy/go-rest-api/entity"
	"github.com/bdxygy/go-rest-api/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PersonControllerImpl struct {
	*service.PersonServiceImpl
}

func (controller *PersonControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	personCreateRequest := entity.PersonCreateOrUpdateRequestEntity{}
	err := decoder.Decode(&personCreateRequest)
	exception.Throw(err)

	personResponseEntity := controller.PersonServiceImpl.Create(r.Context(), personCreateRequest)

	response := entity.ResponseEntity{
		Code:    200,
		Status:  true,
		Data:    personResponseEntity,
		Message: "Create New Person Successfully",
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	exception.Throw(err)

}

func (controller *PersonControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	personUpdateRequest := entity.PersonCreateOrUpdateRequestEntity{}
	err := decoder.Decode(&personUpdateRequest)
	exception.Throw(err)

	personUpdateRequest.UUID = params.ByName("personUUID")

	personResponseEntity := controller.PersonServiceImpl.Create(r.Context(), personUpdateRequest)

	response := entity.ResponseEntity{
		Code:    200,
		Status:  true,
		Data:    personResponseEntity,
		Message: "Update Person Successfully",
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	exception.Throw(err)
}

func (controller *PersonControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	controller.PersonServiceImpl.Delete(r.Context(), params.ByName("personUUID"))

	response := entity.ResponseEntity{
		Code:    200,
		Status:  true,
		Data:    nil,
		Message: "Delete User Successfully",
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	exception.Throw(err)
}

func (controller *PersonControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	finded := controller.PersonServiceImpl.FindById(r.Context(), params.ByName("personUUID"))

	response := entity.ResponseEntity{
		Code:    200,
		Status:  true,
		Data:    finded,
		Message: "Finding User Successfully",
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	exception.Throw(err)
}

func (controller *PersonControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	people := controller.PersonServiceImpl.FindAll(r.Context())

	response := entity.ResponseEntity{
		Code:    200,
		Status:  true,
		Data:    people,
		Message: "Finding All User Successfully",
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	exception.Throw(err)
}

package service

import (
	"context"
	"database/sql"
	"github.com/bdxygy/exception"
	"github.com/bdxygy/go-rest-api/entity"
	"github.com/bdxygy/go-rest-api/helper"
	"github.com/bdxygy/go-rest-api/repository"
)

type PersonServiceImpl struct {
	*sql.DB
	repository.PersonRepositoryImpl
}

func (service *PersonServiceImpl) FindAll(ctx context.Context) []entity.PersonResponseEntity {
	tx, err := service.DB.Begin()
	exception.Throw(err)
	defer helper.DeferCommit(tx)

	people := service.PersonRepositoryImpl.FindAll(ctx, tx)
	responses := []entity.PersonResponseEntity{}

	for i := 0; i < len(people); i++ {

		person := helper.TransformResponse(people[i])

		responses = append(responses, person)
	}

	return responses

}

func (service *PersonServiceImpl) FindById(ctx context.Context, personUUID string) entity.PersonResponseEntity {
	tx, err := service.DB.Begin()
	exception.Throw(err)
	defer helper.DeferCommit(tx)

	responsePerson, err := service.PersonRepositoryImpl.FindById(ctx, tx, personUUID)
	exception.Throw(err)

	return helper.TransformResponse(responsePerson)
}

func (service *PersonServiceImpl) Update(ctx context.Context, personEntity entity.PersonEntity) entity.PersonResponseEntity {
	tx, err := service.DB.Begin()
	exception.Throw(err)
	defer helper.DeferCommit(tx)

	updated := service.PersonRepositoryImpl.Update(ctx, tx, personEntity)

	return helper.TransformResponse(updated)
}

func (service *PersonServiceImpl) Delete(ctx context.Context, personUUID string) {
	tx, err := service.DB.Begin()
	exception.Throw(err)
	defer helper.DeferCommit(tx)

	service.PersonRepositoryImpl.Delete(ctx, tx, personUUID)
}

func (service *PersonServiceImpl) Create(ctx context.Context, personEntity entity.PersonEntity) entity.PersonResponseEntity {
	tx, err := service.DB.Begin()
	exception.Throw(err)
	defer helper.DeferCommit(tx)

	created := service.PersonRepositoryImpl.Create(ctx, tx, personEntity)

	return helper.TransformResponse(created)
}

package service

import (
	"context"
	"database/sql"
	"github.com/bdxygy/exception"
	"github.com/bdxygy/go-rest-api/entity"
	"github.com/bdxygy/go-rest-api/helper"
	"github.com/bdxygy/go-rest-api/repository"
	"github.com/go-playground/validator/v10"
)

type PersonServiceImpl struct {
	*sql.DB
	repository.PersonRepositoryImpl
	*validator.Validate
}

func (service *PersonServiceImpl) FindAll(ctx context.Context) []entity.PersonResponseEntity {
	tx, err := service.DB.Begin()
	exception.Throw(err)
	defer helper.DeferCommit(tx)

	people := service.PersonRepositoryImpl.FindAll(ctx, tx)

	var responses []entity.PersonResponseEntity

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

func (service *PersonServiceImpl) Update(ctx context.Context, personEntity entity.PersonCreateOrUpdateRequestEntity) entity.PersonResponseEntity {
	err := service.Validate.Struct(personEntity)
	exception.Throw(err)

	tx, err := service.DB.Begin()
	exception.Throw(err)
	defer helper.DeferCommit(tx)

	finded, err := service.PersonRepositoryImpl.FindById(ctx, tx, personEntity.UUID)
	exception.Throw(err)

	finded.UUID = personEntity.UUID
	finded.Age = personEntity.Age
	finded.Name = personEntity.Name
	finded.City = personEntity.City
	finded.Address = personEntity.Address

	updated := service.PersonRepositoryImpl.Update(ctx, tx, finded)

	return helper.TransformResponse(updated)
}

func (service *PersonServiceImpl) Delete(ctx context.Context, personUUID string) {
	tx, err := service.DB.Begin()
	exception.Throw(err)
	defer helper.DeferCommit(tx)

	finded, err := service.PersonRepositoryImpl.FindById(ctx, tx, personUUID)
	exception.Throw(err)

	service.PersonRepositoryImpl.Delete(ctx, tx, finded.UUID)
}

func (service *PersonServiceImpl) Create(ctx context.Context, personEntity entity.PersonCreateOrUpdateRequestEntity) entity.PersonResponseEntity {
	err := service.Validate.Struct(personEntity)
	exception.Throw(err)

	tx, err := service.DB.Begin()
	exception.Throw(err)
	defer helper.DeferCommit(tx)

	paramPerson := entity.PersonEntity{
		Name:    personEntity.Name,
		Address: personEntity.Address,
		City:    personEntity.City,
		Age:     personEntity.Age,
	}

	created := service.PersonRepositoryImpl.Create(ctx, tx, paramPerson)

	return helper.TransformResponse(created)
}

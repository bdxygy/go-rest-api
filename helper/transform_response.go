package helper

import "github.com/bdxygy/go-rest-api/entity"

func TransformResponse(personEntity entity.PersonEntity) entity.PersonResponseEntity {
	return entity.PersonResponseEntity{
		UUID:    personEntity.UUID,
		Name:    personEntity.Name,
		Address: personEntity.Address,
		City:    personEntity.City,
		Age:     personEntity.Age,
	}
}

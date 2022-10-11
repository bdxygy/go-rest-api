package service

import (
	"context"
	"github.com/bdxygy/go-rest-api/entity"
)

type PersonService interface {
	FindAll(ctx context.Context) []entity.PersonResponseEntity
	FindById(ctx context.Context, personUUID string) entity.PersonResponseEntity
	Update(ctx context.Context, personEntity entity.PersonCreateOrUpdateRequestEntity) entity.PersonResponseEntity
	Delete(ctx context.Context, personUUID string)
	Create(ctx context.Context, personEntity entity.PersonCreateOrUpdateRequestEntity) entity.PersonResponseEntity
}

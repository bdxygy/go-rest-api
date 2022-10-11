package repository

import (
	"context"
	"database/sql"
	"github.com/bdxygy/go-rest-api/entity"
)

type PersonRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []entity.PersonEntity
	FindById(ctx context.Context, tx *sql.Tx, personUUId string) (entity.PersonEntity, error)
	Update(ctx context.Context, tx *sql.Tx, person entity.PersonEntity) entity.PersonEntity
	Delete(ctx context.Context, tx *sql.Tx, personUUId string)
	Create(ctx context.Context, tx *sql.Tx, person entity.PersonEntity) entity.PersonEntity
}

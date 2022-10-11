package service

import (
	"context"
	"database/sql"
	"github.com/bdxygy/go-rest-api/entity"
)

type PersonServiceImpl struct {
	*sql.DB
}

func (p *PersonServiceImpl) FindAll(ctx context.Context) []entity.PersonResponseEntity {
	//TODO implement me
	panic("implement me")
}

func (p *PersonServiceImpl) FindById(ctx context.Context, personUUID string) entity.PersonResponseEntity {
	//TODO implement me
	panic("implement me")
}

func (p *PersonServiceImpl) Update(ctx context.Context, personEntity entity.PersonEntity) entity.PersonResponseEntity {
	//TODO implement me
	panic("implement me")
}

func (p *PersonServiceImpl) Delete(ctx context.Context, personUUID string) {
	//TODO implement me
	panic("implement me")
}

func (p *PersonServiceImpl) Create(ctx context.Context, personEntity entity.PersonEntity) entity.PersonResponseEntity {
	//TODO implement me
	panic("implement me")
}

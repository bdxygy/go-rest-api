package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/bdxygy/exception"
	"github.com/bdxygy/go-rest-api/entity"
	"github.com/bdxygy/uid/v2"
)

type PersonRepositoryImpl struct {
}

func (p *PersonRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.PersonEntity {

	sqlQuery := fmt.Sprintf("SELECT uuid, name, address, city, age FROM %s", "person")
	res, err := tx.QueryContext(ctx, sqlQuery)

	exception.Throw(err)

	var people []entity.PersonEntity

	for res.Next() {
		person := entity.PersonEntity{}
		err := res.Scan(&person.UUID, &person.Name, &person.Address, &person.City, &person.Age)
		exception.Throw(err)
		people = append(people, person)
	}

	return people

}

func (p *PersonRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, personUUId string) (entity.PersonEntity, error) {
	sqlQuery := fmt.Sprintf("select uuid, name, address, city, age from %s where uuid = ?", "person")

	queryContext, err := tx.QueryContext(ctx, sqlQuery, personUUId)

	exception.Throw(err)

	person := entity.PersonEntity{}
	if queryContext.Next() {
		err := queryContext.Scan(&person.UUID, &person.Name, &person.Address, &person.City, &person.Age)
		exception.Throw(err)
	} else {
		return person, errors.New("Data not found!")
	}

	return person, nil
}

func (p *PersonRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, person entity.PersonEntity) entity.PersonEntity {
	sqlQuery := fmt.Sprintf("UPDATE %s SET name = ?, address = ?, city = ?, age = ? WHERE uuid = ?")
	res, err := tx.ExecContext(ctx, sqlQuery, person.Name, person.Address, person.City, person.Age, person.UUID)
	exception.Throw(err)

	_, err = res.LastInsertId()
	exception.Throw(err)

	return person
}

func (p *PersonRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, personUUId string) {
	sqlQuery := fmt.Sprintf("DELETE FROM %s WHERE uuid = ?", "person")

	_, err := tx.ExecContext(ctx, sqlQuery, personUUId)
	exception.Throw(err)
}

func (p *PersonRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, person entity.PersonEntity) entity.PersonEntity {
	sqlQuery := fmt.Sprintf("INSERT INTO %s (uuid, name, address, city, ages) VALUES (?, ?,?,?,?)", "person")

	res, err := tx.ExecContext(ctx, sqlQuery, uid.Make(), person.Name, person.Address, person.City, person.Age)
	exception.Throw(err)

	id, err := res.LastInsertId()
	exception.Throw(err)

	person.Id = int(id)
	return person
}

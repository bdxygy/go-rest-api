package helper

import (
	"database/sql"
	"github.com/bdxygy/exception"
)

func DeferCommit(tx *sql.Tx) {
	errorMessage := recover()

	if errorMessage != nil {
		err := tx.Rollback()
		exception.Throw(err)
		panic(errorMessage)
	} else {
		err := tx.Commit()
		exception.Throw(err)
	}

}

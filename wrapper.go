package tmpq

import (
	"database/sql"

	// for package initialization.
	_ "github.com/lib/pq"

	"github.com/trimark-jp/errors"
)

type (
	// Wrapper wraps sql.DB.
	Wrapper struct {
		DB *sql.DB
	}
)

const (
	// SQLDriverName is the SQL Driver Name.
	SQLDriverName = "postgres"
)

// Initialize init DB connection.
func (w *Wrapper) Initialize(cs *ConnectionString) error {
	var err error
	w.DB, err = sql.Open(SQLDriverName, cs.String())
	if err != nil {
		return errors.WrapBySourceMsg(err, ErrDBOpenFailed)
	}
	return nil
}

// AutoTx executes the function passing transaction object.
func (w *Wrapper) AutoTx(f TxFunc) (err error) {
	var tx *sql.Tx

	tx, err = w.DB.Begin()
	if err != nil {
		return errors.Wrap(err, ErrBeginTx)
	}

	defer func() {
		if err == nil {
			commitErr := tx.Commit()
			err = errors.Merge(err, errors.Wrap(commitErr, ErrCommit))
		} else {
			rollbackErr := tx.Rollback()
			err = errors.Merge(err, errors.Wrap(rollbackErr, ErrRollback))
		}
	}()

	err = f(tx)
	return err
}

// ExecInsert execute insertion.
func (w *Wrapper) ExecInsert(f InsertFunc) (int, error) {
	var err error
	var resultID int

	err = w.AutoTx(func(tx *sql.Tx) error {
		var inner error
		resultID, inner = f(tx)
		return inner
	})

	return resultID, err
}

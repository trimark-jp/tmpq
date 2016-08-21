package tmpq

import (
	"database/sql"
)

type (
	// TxFunc is function uses transaction.
	TxFunc func(*sql.Tx) error

	// InsertFunc implements insertion.
	// This returns a record ID and error.
	InsertFunc func(*sql.Tx) (int, error)
)

var (
	defaultWrapper = &Wrapper{}
)

// RowsNotFound returns true if the reason of
// the error is that no rows are found.
func RowsNotFound(err error) bool {
	return err == sql.ErrNoRows
}

// Initialize init DB connection.
func Initialize(cs *ConnectionString) error {
	return defaultWrapper.Initialize(cs)
}

// AutoTx executes the function passing transaction object.
func AutoTx(f TxFunc) error {
	return defaultWrapper.AutoTx(f)
}

// ExecInsert execute insertion.
func ExecInsert(f InsertFunc) (int, error) {
	return defaultWrapper.ExecInsert(f)
}

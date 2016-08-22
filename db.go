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

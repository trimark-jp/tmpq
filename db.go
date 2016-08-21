package tmpq

import

// DB Driver for PostgreSQL.
(
	"database/sql"
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
func AutoTx(f func(*sql.Tx) error) error {
	return defaultWrapper.AutoTx(f)
}

package tmpq

import (
	"database/sql"

	// DB Driver for PostgreSQL.
	_ "github.com/lib/pq"
	"github.com/trimark-jp/errors"
)

// Wrapper wraps sql.DB.
type Wrapper struct {
	DB *sql.DB
}

const (
	// SQLDriverName is the SQL Driver Name.
	SQLDriverName = "postgres"
)

// Error Messages.
const (
	ErrDBOpenFailed = "DB open failed"
)

var (
	defaultWrapper = &Wrapper{}
)

// Initialize init DB connection.
func Initialize(cs *ConnectionString) error {
	return defaultWrapper.Initialize(cs)
}

// Initialize init DB connection.
func (w *Wrapper) Initialize(cs *ConnectionString) error {
	var err error
	w.DB, err = sql.Open(SQLDriverName, cs.String())
	if err != nil {
		return errors.WrapBySourceMsg(err, ErrDBOpenFailed)
	}
	return nil
}

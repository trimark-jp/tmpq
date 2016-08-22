package tmpq

import (
	"database/sql"
	"strings"

	"github.com/trimark-jp/errors"
)

// Error Messages.
const (
	ErrDBOpenFailed = "DB open failed"
	ErrBeginTx      = "can't begin transaction"
	ErrCommit       = "can't commit"
	ErrRollback     = "can't rollback"
)

// Error Messages for users.
const (
	ErrTxPrepare          = "Prepare failed"
	ErrQuery              = "Query failed"
	ErrScan               = "Scan failed"
	ErrNotFound           = "Not Found"
	ErrInvalidParamPrefix = "Invalid Parameter: "
	ErrInvalidParamFormat = ErrInvalidParamPrefix + "%s"
)

// RowsNotFound returns true if the reason of
// the error is that no rows are found.
func RowsNotFound(err error) bool {
	return err == sql.ErrNoRows
}

// InvalidParam returns true if the reason of
// the error is that invalid parameter is passed.
func InvalidParam(err error) bool {
	if err == nil {
		return false
	}
	return strings.HasPrefix(errors.SourceOf(err).Error(), ErrInvalidParamPrefix)
}

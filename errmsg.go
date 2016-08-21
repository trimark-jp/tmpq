package tmpq

// Error Messages.
const (
	ErrDBOpenFailed = "DB open failed"
	ErrBeginTx      = "can't begin transaction"
	ErrCommit       = "can't commit"
	ErrRollback     = "can't rollback"
)

// Error Messages for users.
const (
	ErrTxPrepare = "Prepare failed"
	ErrQuery     = "Query failed"
	ErrScan      = "Scan failed"
)

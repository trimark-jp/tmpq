package tmpq

type (
	// RowScanner provides a common interface for sql Scan().
	RowScanner interface {
		Scan(...interface{}) error
	}
)

package tmpq

import (
	"bytes"
	"fmt"
)

type (
	// OffsetLimit implements OFFSET and LIMIT clauses.
	OffsetLimit struct {
		Offset int
		Limit  int
	}
)

const (
	sqlOffsetFormat = ` OFFSET %d `
	sqlLimitFormat  = ` LIMIT %d `
)

// Clause returns OFFSET and LIMIT clause.
func (ol *OffsetLimit) Clause() string {
	buf := new(bytes.Buffer)
	if 0 < ol.Offset {
		fmt.Fprint(buf, sqlOffsetFormat, ol.Offset)
	}
	if 0 < ol.Limit {
		fmt.Fprint(buf, sqlLimitFormat, ol.Limit)
	}
	return buf.String()
}

package datasources

import (
	"github.com/mschoch/tuq/planner"
)

// some utility functions to reduce duplciated code in datasources

func WrapDocWithDatasourceAs(as string, doc planner.Document) planner.Document {
	if as != "" {
		doccopy := doc
		doc = make(planner.Document)
		doc[as] = doccopy
	}
	return doc
}

package optimzer

import (
	"github.com/mschoch/go-unql-couchbase/planner"
)

// the optimizer takes a slice of plans
// tweaks them, and ultimately chooses one for execution

type Optimizer interface {
	Optimize([]planner.Plan) (planner.Plan)
}

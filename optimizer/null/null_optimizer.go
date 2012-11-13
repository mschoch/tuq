package naive

import (
    "github.com/mschoch/go-unql-couchbase/planner"
)

type NullOptimizer struct {
}

func NewNullOptimizer() *NullOptimizer {
    return &NullOptimizer{}
}

func (no *NullOptimizer) Optimize(plans []planner.Plan) planner.Plan {

    // this planner only considers the first option
    plan := plans[0]

    // return the plan
    return plan
}

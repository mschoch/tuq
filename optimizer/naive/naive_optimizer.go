package naive

import (
	"github.com/mschoch/go-unql-couchbase/planner"
	//"log"
)

type NaiveOptimizer struct {
}

func NewNaiveOptimizer() *NaiveOptimizer {
	return &NaiveOptimizer{}
}

func (no *NaiveOptimizer) Optimize(plans []planner.Plan) planner.Plan {

	// this planner only considers the first option
	plan := plans[0]

	// this planner will only try to optimize the plan if it contains 0 joins
	if !PlanContainsJoin(plan) {

		no.MoveWhereToDataSource(plan)

		no.MoveGroupByToDataSource(plan)

		no.MoveOrderByToDataSource(plan)

		no.MoveOffsetToDataSource(plan)

		no.MoveLimitToDataSource(plan)

	}

	// return the plan
	return plan
}

// Support

func (no *NaiveOptimizer) MoveWhereToDataSource(plan planner.Plan) {
	// look for a WHERE clause immediately followed by a data source
	// then ask the datasource if can support this WHERE clause
	// if it can, have it do so, and remove the in memory WHERE
	var last planner.PlanPipelineComponent
	curr := plan.Root.GetSource()
	for curr != nil {
		switch currFilter := curr.(type) {
		case planner.Filter:
			next := curr.GetSource()
			switch nextDataSource := next.(type) {
			case planner.DataSource:
				//log.Printf("Found where clause immidiately followed by datasource %v", last)
				err := nextDataSource.SetFilter(currFilter.GetFilter())
				if err != nil {
					//log.Printf("this datasource cannot support this filter")
				} else {
					//log.Printf("datasource accepted the filter, removing old filter")
					if last == nil {
						plan.Root.SetSource(next)
					} else {
						last.SetSource(next)
					}
				}
			}
		}
		last = curr
		curr = curr.GetSource()
	}
}

func (no *NaiveOptimizer) MoveGroupByToDataSource(plan planner.Plan) {
	// look for a GROUP BY clause immediately followed by a data source
	// then ask the datasource if it can support this GROUP BY clause
	// if it can, have it do so, and remove the in memory GROUP BY
	var last planner.PlanPipelineComponent
	curr := plan.Root.GetSource()
	for curr != nil {
		switch currGrouper := curr.(type) {
		case planner.Grouper:
			next := curr.GetSource()
			switch nextDataSource := next.(type) {
			case planner.DataSource:
				//log.Printf("Found group by clause immidiately followed by datasource %v", last)
				err := nextDataSource.SetGroupByWithStatsFields(currGrouper.GetGroupByWithStatsFields())
				if err != nil {
					//log.Printf("this datasource cannot support this group by")
				} else {
					//log.Printf("datasource accepted the group by, removing old group by")
					if last == nil {
						plan.Root.SetSource(next)
					} else {
						last.SetSource(next)
					}
				}
			}
		}
		last = curr
		curr = curr.GetSource()
	}
}

func (no *NaiveOptimizer) MoveLimitToDataSource(plan planner.Plan) {
	// look for a LIMIT clause immediately followed by a data source
	// then ask the datasource if it can support this LIMIT clause
	// if it can, have it do so, and remove the in memory LIMIT
	var last planner.PlanPipelineComponent
	curr := plan.Root.GetSource()
	for curr != nil {
		switch currLimitter := curr.(type) {
		case planner.Limitter:
			next := curr.GetSource()
			switch nextDataSource := next.(type) {
			case planner.DataSource:
				//log.Printf("Found limit clause immidiately followed by datasource %v", last)
				err := nextDataSource.SetLimit(currLimitter.GetLimit())
				if err != nil {
					//log.Printf("this datasource cannot support this limit")
				} else {
					//log.Printf("datasource accepted the limit, removing old group by")
					if last == nil {
						plan.Root.SetSource(next)
					} else {
						last.SetSource(next)
					}
				}
			}
		}
		last = curr
		curr = curr.GetSource()
	}
}

func (no *NaiveOptimizer) MoveOffsetToDataSource(plan planner.Plan) {
	// look for a OFFSET clause immediately followed by a data source
	// then ask the datasource if it can support this OFFSET clause
	// if it can, have it do so, and remove the in memory OFFSET
	var last planner.PlanPipelineComponent
	curr := plan.Root.GetSource()
	for curr != nil {
		switch currOffsetter := curr.(type) {
		case planner.Offsetter:
			next := curr.GetSource()
			switch nextDataSource := next.(type) {
			case planner.DataSource:
				//log.Printf("Found offset clause immidiately followed by datasource %v", last)
				err := nextDataSource.SetOffset(currOffsetter.GetOffset())
				if err != nil {
					//log.Printf("this datasource cannot support this offset")
				} else {
					//log.Printf("datasource accepted the offset, removing old group by")
					if last == nil {
						plan.Root.SetSource(next)
					} else {
						last.SetSource(next)
					}
				}
			}
		}
		last = curr
		curr = curr.GetSource()
	}
}

func (no *NaiveOptimizer) MoveOrderByToDataSource(plan planner.Plan) {
	// look for a ORDER BY clause immediately followed by a data source
	// then ask the datasource if it can support this ORDER BY clause
	// if it can, have it do so, and remove the in memory ORDER BY
	var last planner.PlanPipelineComponent
	curr := plan.Root.GetSource()
	for curr != nil {
		switch currOrderer := curr.(type) {
		case planner.Orderer:
			next := curr.GetSource()
			switch nextDataSource := next.(type) {
			case planner.DataSource:
				//log.Printf("Found order by clause immidiately followed by datasource %v", last)
				err := nextDataSource.SetOrderBy(currOrderer.GetOrderBy())
				if err != nil {
					//log.Printf("this datasource cannot support this order by")
				} else {
					//log.Printf("datasource accepted the order by, removing old group by")
					if last == nil {
						plan.Root.SetSource(next)
					} else {
						last.SetSource(next)
					}
				}
			}
		}
		last = curr
		curr = curr.GetSource()
	}
}

func PlanContainsJoin(plan planner.Plan) bool {
	component := plan.Root.GetSource()
	for component != nil {
		switch component.(type) {
		case planner.Joiner:
			return true
		}
		// advance
		component = component.GetSource()
	}
	return false
}

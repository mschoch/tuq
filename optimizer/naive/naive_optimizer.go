package naive

import (
	"github.com/mschoch/tuq/parser"
	"github.com/mschoch/tuq/planner"
	"log"
	"strings"
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
		//log.Printf("optimizing non join")
		no.MoveWhereToDataSource(plan)

		no.MoveGroupByToDataSource(plan)

		no.MoveOrderByToDataSource(plan)

		no.MoveOffsetToDataSource(plan)

		no.MoveLimitToDataSource(plan)

	} else {
		//log.Printf("optimizing join")
		// attempting to optimize through a join
		no.MoveWhereToJoin(plan)
		//log.Printf("moved where")
		// see if we can separate any where conditions
		// and apply them directly to the source of the join
		no.MoveJoinConditionsUpTree(plan)
		//log.Printf("moved join condition")
		// now see if the join could be done using sort merge join
		no.TrySortMerge(plan)
		//log.Printf("tried sort merge")

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

func (no *NaiveOptimizer) MoveWhereToJoin(plan planner.Plan) {
	last, curr, next := planner.FindNextPipelineComponentOfTypeFollowedbyType(plan.Root.GetSource(), planner.FilterType, planner.JoinerType)
	if curr != nil && next != nil {
		currFilter := curr.(planner.Filter)
		nextJoiner := next.(planner.Joiner)
		//log.Printf("Found where clause immidiately followed by datasource %v", last)
		err := nextJoiner.SetCondition(currFilter.GetFilter())
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

func (no *NaiveOptimizer) MoveJoinConditionsUpTree(plan planner.Plan) {

	_, curr := planner.FindNextPipelineComponentOfType(plan.Root.GetSource(), planner.JoinerType)

	if curr != nil {
		currJoiner := curr.(planner.Joiner)
		//log.Printf("Expression considerered for separation is %#v", currJoiner.GetCondition())

		// we need to keep running this until it sep is nil
		joinCondition := currJoiner.GetCondition()
		if joinCondition != nil {
			sep, rest := LookForSeparableExpression(joinCondition)
			for sep != nil {

				//log.Printf("Sep Expression is %v", sep)
				//log.Printf("Rest Expression is %v", rest)

				// now try to place the separable expression
				sepSymbols := sep.SymbolsReferenced()
				if len(sepSymbols) > 0 {
					sepDataSource := sepSymbols[0]
					sepDotIndex := strings.Index(sepDataSource, ".")
					sepDs := sepDataSource[0:sepDotIndex]

					// find the datasource on the LHS
					_, leftDs := planner.FindNextPipelineComponentOfType(currJoiner.GetLeftSource(), planner.DataSourceType)
					if leftDs != nil {
						leftDataSource := leftDs.(planner.DataSource)
						if leftDataSource.GetAs() == sepDs {
							//log.Printf("this expression goes left")
							moved := MoveSeparableExpressionToDataSource(leftDataSource, currJoiner, sep, rest)
							if !moved {
								// as soon as we get to something we couldnt move, we stop trying
								return
							}
						}
					}

					// find the datasource on the RHS
					_, rightDs := planner.FindNextPipelineComponentOfType(currJoiner.GetRightSource(), planner.DataSourceType)
					if rightDs != nil {
						rightDataSource := rightDs.(planner.DataSource)
						if rightDataSource.GetAs() == sepDs {
							//log.Printf("this expression goes right")
							moved := MoveSeparableExpressionToDataSource(rightDataSource, currJoiner, sep, rest)
							if !moved {
								// as soon as we get to something we couldnt move, we stop trying
								return
							}
						}
					}

				}

				joinCondition = currJoiner.GetCondition()
				sep, rest = LookForSeparableExpression(joinCondition)
			}
		}
	}

}

func MoveSeparableExpressionToDataSource(dataSource planner.DataSource, joiner planner.Joiner, sep parser.Expression, rest parser.Expression) bool {
	newFilterExpression := sep
	// get the existing filter (if any)
	filterExpression := dataSource.GetFilter()
	if filterExpression != nil {
		// build an AND expression with sep and the existing
		newFilterExpression = parser.NewAndExpression(sep, filterExpression)

	}

	err := dataSource.SetFilter(newFilterExpression)
	if err != nil {
		// this data source doesn't support this filter
		// continue once this becomes a loop
		//log.Printf("Datasource rejected the new filter expression, %v", err)
		return false
	} else {
		// it accepted the filter, now we just updated the join condition
		err := joiner.SetCondition(rest)
		if err != nil {
			// hmmn, now joiner rejected its new condition
			// we must back out the other change
			//log.Printf("Datasource accepted, but now joiner rejected new filter expression, trying to reset")
			err := dataSource.SetFilter(filterExpression)
			if err != nil {
				log.Fatalf("Datasource rejected setting its filter back to the original value, I don't know what to do")
			}
			return false
		}
	}
	return true
}

func LookForSeparableExpression(expr parser.Expression) (parser.Expression, parser.Expression) {
	count := NumberOfDataSourcesReferencedInExpression(expr)
	if count < 2 {
		// all symbols in this expression refer to the same datasource
		// it could be moved out
		return expr, nil
	} else {
		// this expression refers to symbols in 2 or more datasources
		// it cannot be moved itself, but we may have to consier sub-expressions
		and_expr, ok := expr.(*parser.AndExpression)
		if ok {
			// this expression is an AND expression so we can consider moving out
			// its pieces
			sep, remaining := LookForSeparableExpression(and_expr.Left)
			if sep != nil {
				// we found someting to remove
				if remaining == nil {
					// nothing remaining, return RHS
					return sep, and_expr.Right
				} else {
					// build a new AND expression with the remaining piece and the RHS
					return sep, parser.NewAndExpression(remaining, and_expr.Right)
				}
			} else {
				// didn't find anything to remove, try the RHS
				sep, remaining := LookForSeparableExpression(and_expr.Right)
				if sep != nil {
					// we found someting to remove
					if remaining == nil {
						// nothing left, return LHS
						return sep, and_expr.Left
					} else {
						// build a new AND expression with the remaining piece and the RHS
						return sep, parser.NewAndExpression(and_expr.Left, remaining)
					}
				}
			}
		}
	}
	// not an AND expression, nothing more we can do
	return nil, nil
}

func NumberOfDataSourcesReferencedInExpression(expr parser.Expression) int {
	datasourceRef := make(map[string]interface{})
	symbolsReferenced := expr.SymbolsReferenced()
	for _, symbol := range symbolsReferenced {
		datasourceRef[symbol] = nil
	}
	return len(datasourceRef)
}

func (no *NaiveOptimizer) TrySortMerge(plan planner.Plan) {
	last, curr := planner.FindNextPipelineComponentOfType(plan.Root.GetSource(), planner.JoinerType)
	if curr != nil {
		currJoiner := curr.(planner.Joiner)
		sortMergeJoiner := planner.NewOttoSortMergeJoiner()
		err := sortMergeJoiner.SetCondition(currJoiner.GetCondition())
		if err == nil {
			// now we need to make sure that the upstream sources can be sorted propertly
			leftSource := currJoiner.GetLeftSource()
			rightSource := currJoiner.GetRightSource()

			// FIXME for now this only works if the immidiate sources are datasources
			// needs to be made to work through joins and other pipeline components

			leftDataSource, isLeftDataSource := leftSource.(planner.DataSource)
			rightDataSource, isRightDataSource := rightSource.(planner.DataSource)
			if isLeftDataSource && isRightDataSource {
				if leftDataSource.GetOrderBy() == nil && rightDataSource.GetOrderBy() == nil {
					leftSort := parser.NewSortItem(parser.NewProperty(sortMergeJoiner.LeftExpr.SymbolsReferenced()[0]), true)
					err = leftDataSource.SetOrderBy(parser.SortList{*leftSort})
					if err != nil {
						//log.Printf("merge join not possible left datasource rejected order by")
						return
					}

					rightSort := parser.NewSortItem(parser.NewProperty(sortMergeJoiner.RightExpr.SymbolsReferenced()[0]), true)
					err = rightDataSource.SetOrderBy(parser.SortList{*rightSort})
					if err != nil {
						//log.Printf("merge join not possible right datasource rejected order by")
						return
					}

					//if we made it this far, it shoudl work
					//lets attach the left and right datasources to the new joiner
					sortMergeJoiner.SetLeftSource(leftDataSource)
					sortMergeJoiner.SetRightSource(rightDataSource)

					//now we just need to replace the existing joiner with the new one 
					if last == nil {
						plan.Root.SetSource(sortMergeJoiner)
					} else {
						last.SetSource(sortMergeJoiner)
					}
				} else {
					//log.Printf("merge join optimization only supports direct datasources that are not already ordered")
				}
			} else {
				//log.Printf("merge join optimization only supports direct datasources before join")
			}
		} else {
			//log.Printf("merge join not going to work here %v", err)
		}
	}
}

package naive

import (
	"github.com/mschoch/tuq/datasources"
	"github.com/mschoch/tuq/parser"
	"github.com/mschoch/tuq/planner"
	"log"
)

// this is the simplest planner possible
// it  will join data sources in the order it finds them
// and it will perform all operations in memory

// sensible behavior is provided by the optimizer
// this is just trying to produce correct results

// FIXME future versions should at least produce multiple plans
// representing different join orders
// the optimizer can then choose between the alternatives

type NaivePlanner struct {
}

func NewNaivePlanner() *NaivePlanner {
	return &NaivePlanner{}
}

func (np *NaivePlanner) Plan(query parser.Select) []planner.Plan {
	result := planner.Plan{}

	var last_in_pipeline planner.PlanPipelineComponent

	// from
	for _, datasource := range query.From {

		// look up the data source		
		ds := datasources.NewDataSourceWithName(datasource.Def)
		if ds == nil {
			log.Printf("No such datasource exists")
			return nil
		}

		ds.SetName(datasource.Def)
		ds.SetAs(datasource.As)

		if last_in_pipeline != nil {
			joiner := planner.NewOttoCartesianProductJoiner()
			joiner.SetLeftSource(last_in_pipeline)
			joiner.SetRightSource(ds)
			last_in_pipeline = joiner
		} else {
			last_in_pipeline = ds
		}
	}

	// where
	if query.Where != nil {
		ottoFilter := planner.NewOttoFilter()
		ottoFilter.SetSource(last_in_pipeline)
		ottoFilter.SetFilter(query.Where)
		last_in_pipeline = ottoFilter
	}

	if query.IsAggregateQuery() {
		// in an aggregate query we sometimes need to collect stats
		// on fields that we're not aggregating on
		// the full list will be the symbols from the SELECT clause
		// invalid entries should have been detected earlier 
		// (columns not in the group by clause and not inside an aggregate function)
		stat_fields := make([]string, 0)
		if query.Sel != nil {
			stat_fields = query.Sel.SybolsReferenced()
		}

		// group by
		if query.Groupby != nil {
			ottoGrouper := planner.NewOttoGrouper()
			ottoGrouper.SetGroupByWithStatsFields(query.Groupby, stat_fields)
			ottoGrouper.SetSource(last_in_pipeline)
			last_in_pipeline = ottoGrouper
		} else {
			// easy to overlook, but if they use aggregate functions
			// this is aggregate query, even though they didnt explicitly
			// have a group by clause, we should apply a default group by
			// clause that groups everything
			ottoGrouper := planner.NewOttoGrouper()
			ottoGrouper.SetGroupByWithStatsFields(parser.ExpressionList{parser.NewBoolLiteral(true)}, stat_fields)
			ottoGrouper.SetSource(last_in_pipeline)
			last_in_pipeline = ottoGrouper
		}

		// having
		if query.Having != nil {
			ottoFilter := planner.NewOttoFilter()
			ottoFilter.SetSource(last_in_pipeline)
			ottoFilter.SetFilter(query.Having)
			last_in_pipeline = ottoFilter
		}
	}

	// order by
	if query.Orderby != nil {
		ottoOrderer := planner.NewOttoOrderer()
		ottoOrderer.SetSource(last_in_pipeline)
		ottoOrderer.SetOrderBy(query.Orderby)
		last_in_pipeline = ottoOrderer
	}

	// offset
	if query.Offset != nil {
		offset := planner.NewOttoOffsetter()
		offset.SetOffset(query.Offset)
		offset.SetSource(last_in_pipeline)
		last_in_pipeline = offset
	}

	// limit
	if query.Limit != nil {
		limit := planner.NewOttoLimitter()
		limit.SetLimit(query.Limit)
		limit.SetSource(last_in_pipeline)
		last_in_pipeline = limit
	}

	// select

	if query.Sel != nil {
		ottoSelecter := planner.NewOttoSelecter()
		ottoSelecter.SetSource(last_in_pipeline)
		ottoSelecter.SetSelect(query.Sel)
		result.Root = ottoSelecter
	} else {
		defaultSelecter := planner.NewDefaultSelecter()
		defaultSelecter.SetSource(last_in_pipeline)
		result.Root = defaultSelecter
	}

	return []planner.Plan{result}

}

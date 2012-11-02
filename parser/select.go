package parser

import ()

type Select struct {
	Distinct           bool
	Sel                Expression
	SelAs              string
	From               []DataSource
	Where              Expression
	Groupby            ExpressionList
	Having             Expression
	Orderby            SortList
	Limit              Expression
	Offset             Expression
	parsedSuccessfully bool
	isAggregateQuery   bool
}

func NewSelect() *Select {
	return &Select{
		Distinct:           false,
		parsedSuccessfully: false,
		isAggregateQuery:   false}
}

func (s *Select) AddDataSource(ds *DataSource) {
	s.From = append(s.From, *ds)
}

func (s *Select) WasParsedSuccessfully() bool {
	return s.parsedSuccessfully
}

func (s *Select) IsAggregateQuery() bool {
	return s.isAggregateQuery
}

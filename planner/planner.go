package planner

import (
	"github.com/mschoch/go-unql-couchbase/parser"
)

type Document map[string]interface{}
type DocumentChannel chan Document

type Row interface{}
type RowChannel chan Row

// the planner takes a parsed select query
// and returns a set of one or more plans
// later, an optimizer can tweak them, and select one to execute

type Planner interface {
	Plan(parser.Select, map[string]interface{}) []Plan
}

// a plan is the root of a tree
// the tree is composed of plan pipeline components
// the leaves of the tree will always be a data source
type Plan struct {
	Root Selecter
}

func (p *Plan) Run() RowChannel {
	go p.Root.Run()
	rowChannel := p.Root.GetRowChannel()
	return rowChannel
}

func (p *Plan) Explain() RowChannel {
	go p.Root.Explain()
	rowChannel := p.Root.GetRowChannel()
	return rowChannel
}

type PlanPipelineComponent interface {
	SetSource(PlanPipelineComponent)
	GetSource() PlanPipelineComponent
	GetDocumentChannel() DocumentChannel
	Run()
	Explain()
	Cancel() // there are times that the downstream component knows it doesn't need any more from the upstream
}

type Offsetter interface {
	SetOffset(parser.Expression)
	GetOffset() parser.Expression
}

type Limitter interface {
	SetLimit(parser.Expression)
	GetLimit() parser.Expression
}

type Filter interface {
	SetFilter(parser.Expression)
	GetFilter() parser.Expression
}

type Grouper interface {
	SetGroupByWithStatsFields(parser.ExpressionList, []string)
    GetGroupByWithStatsFields() (parser.ExpressionList, []string)
}

type Orderer interface {
	SetOrderBy(parser.SortList)
	GetOrderBy() parser.SortList
}

type Joiner interface {
	SetSource(PlanPipelineComponent)
	GetSource() PlanPipelineComponent
	GetDocumentChannel() DocumentChannel
	Run()
	Explain()
	Cancel()
	SetLeftSource(PlanPipelineComponent)
	SetRightSource(PlanPipelineComponent)
	SetCondition(parser.Expression) error
}

type DataSource interface {
	SetName(string)
	SetAs(string)
	SetFilter(parser.Expression) error
	SetOrderBy(parser.SortList) error
	SetLimit(parser.Expression) error
	SetOffset(parser.Expression) error
	SetGroupByWithStatsFields(parser.ExpressionList, []string) error
	SetHaving(parser.Expression) error
	Cancel()
	Explain()
	GetDocumentChannel() DocumentChannel
	Run()
	SetSource(PlanPipelineComponent)
	GetSource() PlanPipelineComponent
	// FIXME this last method is a bit of a hack
	// it should be handled byy proper support of an IN clause
	DocsFromIds(docIds []string) ([]Document, error)
}

type Selecter interface {
	SetSource(PlanPipelineComponent)
	GetSource() PlanPipelineComponent
	GetRowChannel() RowChannel
	Run()
	Explain()
	Cancel() // there are times that the downstream component knows it doesn't need any more from the upstream
	SetSelect(parser.Expression)
}

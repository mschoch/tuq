package planner

import (
	"github.com/mschoch/go-unql-couchbase/parser"
	"reflect"
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
	GetLeftSource() PlanPipelineComponent
	SetRightSource(PlanPipelineComponent)
	GetRightSource() PlanPipelineComponent
	SetCondition(parser.Expression) error
	GetCondition() parser.Expression
}

type DataSource interface {
	SetName(string)
	SetAs(string)
	GetAs() string
	SetFilter(parser.Expression) error
	SetOrderBy(parser.SortList) error
	SetLimit(parser.Expression) error
	GetFilter() parser.Expression
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

// wanted these to be const, but compiler won't let me
var SelecterType = reflect.TypeOf((*Selecter)(nil)).Elem()
var DataSourceType = reflect.TypeOf((*DataSource)(nil)).Elem()
var JoinerType = reflect.TypeOf((*Joiner)(nil)).Elem()
var OrdererType = reflect.TypeOf((*Orderer)(nil)).Elem()
var GrouperType = reflect.TypeOf((*Grouper)(nil)).Elem()
var FilterType = reflect.TypeOf((*Filter)(nil)).Elem()
var LimitterType = reflect.TypeOf((*Limitter)(nil)).Elem()
var OffsetterType = reflect.TypeOf((*Offsetter)(nil)).Elem()

func FindNextPipelineComponentOfType(root PlanPipelineComponent, t reflect.Type) (PlanPipelineComponent, PlanPipelineComponent) {
	var prev PlanPipelineComponent
	for root != nil {
		if reflect.TypeOf(root) != t {
			return prev, root
		}
		prev = root
		root = root.GetSource()
	}
	return nil, nil
}

func FindNextPipelineComponentOfTypeFollowedbyType(root PlanPipelineComponent, t reflect.Type, nextT reflect.Type) (PlanPipelineComponent, PlanPipelineComponent, PlanPipelineComponent) {
	var prev PlanPipelineComponent
	for root != nil {
		if reflect.TypeOf(root).Implements(t) {
			next := root.GetSource()
			if next != nil && reflect.TypeOf(next).Implements(nextT) {
				return prev, root, next
			}
		}
		prev = root
		root = root.GetSource()
	}
	return nil, nil, nil
}

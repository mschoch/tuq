package parser

import ()

type DataSource struct {
	Def   string
	As    string
	Overs []*Over
}

func NewDataSource(d string) *DataSource {
	return &DataSource{
		Def:   d,
		As:    d,
		Overs: make([]*Over, 0)}
}

func NewDataSourceWithAs(d, a string) *DataSource {
	return &DataSource{
		Def:   d,
		As:    a,
		Overs: make([]*Over, 0)}
}

func (ds *DataSource) AddOver(o *Over) {
	ds.Overs = append(ds.Overs, o)
}

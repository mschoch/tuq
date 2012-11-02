package parser

import ()

type DataSource struct {
	Def string
	As  string
}

func NewDataSource(d string) *DataSource {
	return &DataSource{
		Def: d,
		As:  d}
}

func NewDataSourceWithAs(d, a string) *DataSource {
	return &DataSource{
		Def: d,
		As:  a}
}

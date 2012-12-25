package parser

import (
	"fmt"
	"strings"
)

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
	isExplainOnly      bool
}

func NewSelect() *Select {
	return &Select{
		Distinct:           false,
		parsedSuccessfully: false,
		isAggregateQuery:   false,
		isExplainOnly:      false}
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

func (s *Select) IsExplainOnly() bool {
	return s.isExplainOnly
}

// NOTE: this should not be used to enforce limitations
//       elsewhere in the system, it should only detect
//       semantic problems in the query itself
func (s *Select) Validate() error {

	if len(s.From) == 0 {
		return fmt.Errorf("Please provide at least one datasource")
	}

	// datasources must be identifiable by different names
	// (either explicitly using AS, or implicitly because they have different names)
	dataSourceByName := make(map[string]interface{})
	for _, v := range s.From {

		// first check the datasource AS
		_, exists := dataSourceByName[v.As]
		if exists {
			return fmt.Errorf("Ambiguous data source names, use AS clause to disambiguate")
		} else {
			dataSourceByName[v.As] = v
		}

		// now check any datasource OVER ... AS
		for _, ov := range v.Overs {
			_, oexists := dataSourceByName[ov.As]
			if oexists {
				return fmt.Errorf("Ambiguous data source names, use AS clause to disambiguate")
			} else {
				dataSourceByName[ov.As] = ov
			}
		}
	}

	// if there is only 1 data source (each OVER AS counts here as well)
	// then the query can omit the datasource name in symbols
	// however, we will add it back internally for our own sanity

	// FIXME reading this code now, does it even allow you to
	// manually specify the prefix, or would we prepend it a 2nd time
	// no matter what?
	if len(dataSourceByName) == 1 {
		prefix := fmt.Sprintf("%v.", s.From[0].As)

		if s.Sel != nil {
			s.Sel.PrefixSymbols(prefix)
		}

		if s.Where != nil {
			s.Where.PrefixSymbols(prefix)
		}
		if s.Having != nil {
			s.Having.PrefixSymbols(prefix)
		}
		if s.Limit != nil {
			s.Limit.PrefixSymbols(prefix)
		}
		if s.Offset != nil {
			s.Offset.PrefixSymbols(prefix)
		}

		if s.Groupby != nil {
			for _, groupby := range s.Groupby {
				groupby.PrefixSymbols(prefix)
			}
		}

		if s.Orderby != nil {
			for _, order := range s.Orderby {
				order.Sort.PrefixSymbols(prefix)
			}
		}
	}

	// at this point, all symbols should now start with a prefix that 
	// identifies the data source, check this now
	err := verifySymbolsValid(s.Sel, dataSourceByName)
	if err != nil {
		return err
	}
	err = verifySymbolsValid(s.Where, dataSourceByName)
	if err != nil {
		return err
	}
	err = verifySymbolsValid(s.Having, dataSourceByName)
	if err != nil {
		return err
	}
	err = verifySymbolsValid(s.Limit, dataSourceByName)
	if err != nil {
		return err
	}
	err = verifySymbolsValid(s.Offset, dataSourceByName)
	if err != nil {
		return err
	}

	if s.Groupby != nil {
		for _, groupby := range s.Groupby {
			err = verifySymbolsValid(groupby, dataSourceByName)
			if err != nil {
				return err
			}
		}
	}

	if s.Orderby != nil {
		for _, order := range s.Orderby {
			err = verifySymbolsValid(order.Sort, dataSourceByName)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func verifySymbolsValid(e Expression, d map[string]interface{}) error {
	if e != nil {
		exprSymbols := e.SymbolsReferenced()
	OUTER:
		for _, v := range exprSymbols {
			for prefixFrom, _ := range d {
				if strings.HasPrefix(v, prefixFrom+".") {
					// this symbol is OK
					continue OUTER
				}
			}
			// if we get here, we have symbol which does not refer to a data source
			return fmt.Errorf("Symbol %v does not refer to a datasource in this query", v)
		}
	}
	return nil
}

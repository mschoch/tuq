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
	dataSourceByName := make(map[string]DataSource)
	for _, v := range s.From {
		_, exists := dataSourceByName[v.As]
		if exists {
			return fmt.Errorf("Ambiguous data source names, use AS clause to disambiguate")
		} else {
			dataSourceByName[v.As] = v
		}
	}

	// if there is only 1 data source, query can omit the datasource name in symbols
	// however, we will add it back internally for our own sanity
	if len(s.From) == 1 {
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
	err := verifySymbolsValid(s.Sel, s.From)
	if err != nil {
		return err
	}
	err = verifySymbolsValid(s.Where, s.From)
	if err != nil {
		return err
	}
	err = verifySymbolsValid(s.Having, s.From)
	if err != nil {
		return err
	}
	err = verifySymbolsValid(s.Limit, s.From)
	if err != nil {
		return err
	}
	err = verifySymbolsValid(s.Offset, s.From)
	if err != nil {
		return err
	}

	if s.Groupby != nil {
		for _, groupby := range s.Groupby {
			err = verifySymbolsValid(groupby, s.From)
			if err != nil {
				return err
			}
		}
	}

	if s.Orderby != nil {
		for _, order := range s.Orderby {
			err = verifySymbolsValid(order.Sort, s.From)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func verifySymbolsValid(e Expression, d []DataSource) error {
	if e != nil {
		exprSymbols := e.SybolsReferenced()
	OUTER:
		for _, v := range exprSymbols {
			for _, prefixFrom := range d {
				if strings.HasPrefix(v, prefixFrom.As+".") {
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

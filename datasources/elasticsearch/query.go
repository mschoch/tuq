package elasticsearch

type QueryDocument map[string]interface{}
type Query map[string]interface{}
type Filter map[string]interface{}
type Facets map[string]interface{}
type Facet map[string]interface{}


func MatchAllQuery() *Query {
	query := make(Query)
	query["match_all"] = map[string]interface{}{}
	return &query
}

func EmptyFilter() *Filter {
	filter := make(Filter)
	return &filter
}

func NewTermFilter(field string, value interface{}) *Filter {
	filter := make(Filter)
	termFilter := make(Filter)
	termFilter[field] = value
	filter["term"] = termFilter
	return &filter
}

func NewNotFilter(operand *Filter) *Filter {
	filter := make(Filter)
	filter["not"] = operand
	return &filter
}

func NewRangeFilter(field string, value interface{}, comparator string) *Filter {
	filter := make(Filter)
	inner := make(Filter)
	rang := make(Filter)
	rang[comparator] = value
	inner[field] = rang
	filter["range"] = inner
	return &filter
}

func NewAndFilter(lhs *Filter, rhs *Filter) *Filter {
	filter := make(Filter)
	and := []Filter{*lhs, *rhs}
	filter["and"] = and
	return &filter
}

func NewOrFilter(lhs *Filter, rhs *Filter) *Filter {
	filter := make(Filter)
	or := []Filter{*lhs, *rhs}
	filter["or"] = or
	return &filter
}

func EmptyFacets() *Facets {
	facets := make(Facets)
	return &facets
}

func NewStatisticalFacet(field string, size int) *Facet {
	facet := make(Facet)
	stat := make(Facet)
	stat["field"] = field
	stat["size"] = size
	facet["statistical"] = stat
	return &facet
}

func NewTermsStatsFacet(keyfield string, valuefield string, size int) *Facet {
	facet := make(Facet)
	stat := make(Facet)
	stat["key_field"] = keyfield
	stat["value_field"] = valuefield
	stat["size"] = size
	facet["terms_stats"] = stat
	return &facet
}

func NewTermsFacet(field string, size int) *Facet {
	facet := make(Facet)
	stat := make(Facet)
	stat["field"] = field
	stat["size"] = size
	facet["terms"] = stat
	return &facet
}

func NewQueryDocument(query *Query, filter *Filter, facets *Facets, size int) *QueryDocument {
	doc := make(QueryDocument)
	doc["query"] = query
	doc["filter"] = filter
	doc["facets"] = facets
	doc["size"] = size
	return &doc
}

func NewDefaultQuery() *QueryDocument {
	return NewQueryDocument(MatchAllQuery(), EmptyFilter(), EmptyFacets(), EsBatchSize)
}

func (f *Facet) SetFacetSize(size int) {
	for k, _ := range *f {
		facet, ok := (*f)[k].(Facet)
		if ok {
			facet["size"] = size
		}
	}
}

func (f *Facet) SetFacetFilter(filter *Filter) {
	(*f)["facet_filter"] = *filter
}


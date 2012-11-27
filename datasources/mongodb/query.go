package mongodb

type Filter map[string]interface{}
type SortItem map[string]interface{}

func EmptyFilter() *Filter {
	filter := make(Filter)
	return &filter
}

func NewEqualsFilter(field string, value interface{}) *Filter {
	filter := make(Filter)
	filter[field] = value
	return &filter
}

func NewNotFilter(operand *Filter) *Filter {
	filter := make(Filter)
	filter["$not"] = operand
	return &filter
}

func NewRangeFilter(field string, value interface{}, comparator string) *Filter {
	filter := make(Filter)
	rang := make(Filter)
	rang[comparator] = value
	filter[field] = rang
	return &filter
}

func NewAndFilter(lhs *Filter, rhs *Filter) *Filter {
	filter := make(Filter)
	and := []Filter{*lhs, *rhs}
	filter["$and"] = and
	return &filter
}

func NewOrFilter(lhs *Filter, rhs *Filter) *Filter {
	filter := make(Filter)
	or := []Filter{*lhs, *rhs}
	filter["$or"] = or
	return &filter
}

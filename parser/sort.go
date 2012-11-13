package parser

import (
	"fmt"
)

type SortItem struct {
	Sort      Expression
	Ascending bool
}

type SortList []SortItem

func NewSortItem(s Expression, a bool) *SortItem {
	return &SortItem{
		Sort:      s,
		Ascending: a}
}

func (si *SortItem) String() string {
	if si.Ascending {
		return fmt.Sprintf("%s ASC", si.Sort)
	}
	return fmt.Sprintf("%s DESC", si.Sort)
}

func (sl SortList) String() string {
	result := ""
	for _, v := range sl {
		result += fmt.Sprintf("%v", v)
	}
	return result
}

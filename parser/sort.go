package parser

import()

type SortItem struct {
    Sort  Expression
    Ascending bool
}

type SortList []SortItem

func NewSortItem(s Expression, a bool) *SortItem {
    return &SortItem{
        Sort:  s,
        Ascending: a}
}
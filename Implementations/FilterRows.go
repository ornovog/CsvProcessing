package Implementations

import prc "csvProcessing/Processing"

type FilterRows struct {
	*prc.MapOperation
	column int
	value  string
}

func NewFilterRows(column int, value string) prc.Operation {
	filterRows := func(row prc.Row, _ int) *prc.Row{
		if len(row) > column && row[column] == value{
			return &row
		}

		return nil
	}
	m := prc.NewMapOperation(filterRows)
	f := &FilterRows{MapOperation: m, column: column, value: value}
	m.Op = f
	return f
}

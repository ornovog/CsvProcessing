package Implementations

import prc "csvProcessing/Processing"

type FilterRows struct {
	*prc.BaseOperation
	column int
	value  string
}

func (f *FilterRows) Run() prc.LazyTable {
	go func() {
		defer close(f.OutputTable)
		for row := range f.InputTable {
			if len(row) > f.column && row[f.column] == f.value {
				f.OutputTable <- row
			}
		}
	}()

	return f.OutputTable
}

func NewFilterRows(column int, value string) prc.Operation {
	base := &prc.BaseOperation{}
	f := &FilterRows{BaseOperation: base, column: column, value: value}
	base.Op = f
	return f
}

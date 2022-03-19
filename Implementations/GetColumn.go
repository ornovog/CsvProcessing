package Implementations

import prc "csvProcessing/Processing"

type GetColumn struct {
	*prc.MapOperation
	column int
}

func NewGetColumn(column int) prc.Operation {
	getColumn := func(row prc.Row, _ int) *prc.Row{
		if len(row) > column{
			return &prc.Row{row[column]}
		}

		return nil
	}
	m := prc.NewMapOperation(getColumn)
	g := &GetColumn{MapOperation: m, column: column}
	m.Op = g
	return g
}

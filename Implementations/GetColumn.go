package Implementations

import prc "csvProcessing/Processing"

type GetColumn struct {
	*prc.MapOperation
	column int
}

func (g *GetColumn) Run() prc.LazyTable {
	go func() {
		defer close(g.OutputTable)
		for row := range g.InputTable {
			if len(row) > g.column {
				g.OutputTable <- prc.Row{row[g.column]}
			}
		}
	}()

	return g.OutputTable
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

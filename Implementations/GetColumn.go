package Implementations

import prc "csvProcessing/Processing"

type GetColumn struct {
	*prc.BaseOperation
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
	base := &prc.BaseOperation{}
	g := &GetColumn{BaseOperation: base, column: column}
	base.Op = g
	return g
}

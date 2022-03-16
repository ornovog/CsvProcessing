package Implementations

import (
	prc "csvProcessing/Processing"
	"fmt"
	"strconv"
)

type Avg struct {
	*prc.BaseOperation
}

func (a *Avg) Run() prc.LazyTable {
	go func() {
		defer close(a.OutputTable)
		var sum, n float32
		for row := range a.InputTable {
			if len(row) == 1 {
				val, err := strconv.Atoi(row[0])
				if err == nil {
					sum += float32(val)
					n++
				}
			}

		}
		if n > 0 {
			avgStr := fmt.Sprintf("%f", sum/n)
			a.OutputTable <- prc.Row{avgStr}
		}
	}()

	return a.OutputTable
}

func NewAvg() prc.Operation {
	base := &prc.BaseOperation{}
	a := &Avg{BaseOperation: base}
	base.Op = a
	return a
}

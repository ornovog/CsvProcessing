package Implementations

import (
	prc "csvProcessing/Processing"
	"math"
	"strconv"
)

type Ceil struct {
	*prc.BaseOperation
}

func (c *Ceil) Run() prc.LazyTable {
	go func() {
		defer close(c.OutputTable)
		for row := range c.InputTable {
			if len(row) == 1 {
				val, err := strconv.ParseFloat(row[0], 64)
				if err == nil {
					ceil := int(math.Ceil(val))
					ceilStr := strconv.Itoa(ceil)
					c.OutputTable <- prc.Row{ceilStr}
				}
			}

		}
	}()

	return c.OutputTable
}

func NewCeil() prc.Operation {
	base := &prc.BaseOperation{}
	c := &Ceil{BaseOperation: base}
	base.Op = c
	return c
}

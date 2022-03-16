package Implementations

import (
	prc "csvProcessing/Processing"
	"math"
	"strconv"
)

type Ceil struct {
	*prc.MapOperation
}

func NewCeil() prc.Operation {
	ceil := func(row prc.Row, _ int) *prc.Row{
		if len(row) == 1{
			val, err := strconv.ParseFloat(row[0], 64)
			if err == nil {
				ceil := int(math.Ceil(val))
				ceilStr := strconv.Itoa(ceil)
				return &prc.Row{ceilStr}
			}
		}
		return nil
	}

	m := prc.NewMapOperation(ceil)
	f := &Ceil{MapOperation: m}
	m.Op = f
	return f
}

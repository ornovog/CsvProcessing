package Implementations

import (
	prc "csvProcessing/Processing"
	"fmt"
	"strconv"
)

type Avg struct {
	*prc.ReduceOperation
}

func NewAvg() prc.Operation {
	var sum *float32
	Sum := func(row prc.Row, index int){
		if len(row) == 1 {
			val, err := strconv.Atoi(row[0])
			if err == nil {
				if sum == nil{
					sum = new(float32)
				}

				*sum += float32(val)
			}
		}
	}

	Div := func(count int) *prc.Row{
		return &prc.Row{fmt.Sprintf("%f", *sum/float32(count))}
	}

	r := prc.NewReduceOperation(Sum, Div)
	a := &Avg{ReduceOperation: r}
	r.Op = a
	return a
}

package Processing

type Reduce func(row Row, index int)

type FinalReduce func(int) *Row

type ReduceOperation struct {
	finalReduce FinalReduce
	iterate     Reduce
	*BaseOperation
}

func (r *ReduceOperation) Run() LazyTable {
	go func() {
		defer close(r.OutputTable)
		index := 0
		for row := range r.InputTable {
			r.iterate(row, index)
			index++
		}

		result := r.finalReduce(index)
		if result != nil {
			r.OutputTable <- *result
		}
	}()

	return r.OutputTable
}

func NewReduceOperation(iterate Reduce, finalReduce FinalReduce) *ReduceOperation {
	base := &BaseOperation{}
	r := &ReduceOperation{BaseOperation: base, iterate: iterate, finalReduce: finalReduce}
	base.Op = r
	return r
}

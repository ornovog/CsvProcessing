package Processing

type BaseOperation struct {
	InputTable, OutputTable LazyTable
	Op                      Operation
}

func (b *BaseOperation) With(nextOperation Operation) Operation {
	nextOperation.Init(b.Op.Run())
	return nextOperation
}

func (b *BaseOperation) Init(inputTable LazyTable) {
	b.OutputTable = make(chan Row)
	b.InputTable = inputTable
}

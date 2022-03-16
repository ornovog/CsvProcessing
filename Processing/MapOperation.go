package Processing

type Map func(row Row, index int) *Row

type MapOperation struct {
	iterate Map
	*BaseOperation
}

func (m *MapOperation) Run() LazyTable {
	go func() {
		defer close(m.OutputTable)
		var result *Row
		index := 0
		for row := range m.InputTable {
			result = m.iterate(row, index)
			if result != nil {
				m.OutputTable <- *result
			}
			index++
		}
	}()

	return m.OutputTable
}

func NewMapOperation(iterate Map) *MapOperation {
	base := &BaseOperation{}
	m := &MapOperation{BaseOperation: base, iterate: iterate}
	base.Op = m
	return m
}

package Processing

type Row []string

type LazyTable chan Row

type Operation interface {
	Init(LazyTable)
	With(nextOperation Operation) Operation
	Run() LazyTable
}

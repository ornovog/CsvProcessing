package Processing

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type CsvLoader struct {
	path string
}

func (c *CsvLoader) Run() LazyTable {
	f, err := os.Open(c.path)
	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(f)
	outTable := make(chan Row)

	go func() {
		defer f.Close()
		defer close(outTable)

		for record, err := csvReader.Read(); err == nil; record, err = csvReader.Read() {
			outTable <- record
		}

		if !(err == nil || err == io.EOF) {
			log.Fatal(err)
		}
	}()

	return outTable
}

func (c *CsvLoader) With(nextOperation Operation) Operation {
	nextOperation.Init(c.Run())
	return nextOperation
}

func NewCsvLoader(path string) *CsvLoader {
	var c CsvLoader
	c.path = path
	return &c
}

package main

import (
	imp "csvProcessing/Implementations"
	"csvProcessing/Processing"
	"fmt"
)


func main() {
	table := Processing.NewCsvLoader("CsvToProcess.csv").
		With(imp.NewFilterRows(3, "Iowa")).
		With(imp.NewGetColumn(10)).
		With(imp.NewAvg()).
		With(imp.NewCeil()).Run()

	for row := range table{
		fmt.Println(row)
	}
}

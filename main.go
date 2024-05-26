package main

import (
	. "main/operations"
	"strconv"
)

func main() {
	ReadCSV("inputLarge.csv").
		With(FilterRows(func(row []string) bool {
			cell, _ := strconv.Atoi(row[0])
			return cell > 100 && cell < 200
		})).
		With(GetColumns(3, 5)).
		With(ConsoleLog()).
		With(Average()).
		With(ConsoleLog()).
		WriteCSV("output.csv")
}

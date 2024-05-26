package operations

import (
	"encoding/csv"
	"fmt"
	"os"
)

// WriteCSV takes a DataStream and writes it into a CSV file with the given filePath.
func (ds *DataStream) WriteCSV(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	for record := range ds.data {
		if err := csvWriter.Write(record); err != nil {
			fmt.Println(err)
			return
		}
	}
}

package operations

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

// ReadCSV reads a CSV file with the given filePath and returns a DataStream
func ReadCSV(filePath string) *DataStream {
	output := make(chan []string)
	file, err := os.Open(filePath)
	if err != nil {

	}

	go func() {
		defer close(output)
		csvReader := csv.NewReader(file)
		for {
			record, err := csvReader.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			output <- record
		}
	}()
	return &DataStream{data: output}
}

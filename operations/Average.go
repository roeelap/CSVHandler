package operations

import (
	"fmt"
	"strconv"
)

// AverageOperation calculates the average of the data
type AverageOperation struct {
	sum   float64
	count float64
}

func (a *AverageOperation) Apply(ds *DataStream) *DataStream {
	output := make(chan []string)
	go func() {
		defer close(output)
		for record := range ds.data {
			for _, value := range record {
				valueAsFloat, err := strconv.ParseFloat(value, 64)
				if err != nil {
					fmt.Println("Value is not a number - " + value)
				}
				a.sum += valueAsFloat
				a.count++
			}
		}
		output <- []string{fmt.Sprintf("%f", a.sum/a.count)}
	}()

	return &DataStream{data: output}
}

func Average() *AverageOperation {
	return &AverageOperation{}
}

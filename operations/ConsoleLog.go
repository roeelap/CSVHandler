package operations

import (
	"fmt"
)

// ConsoleLogOperation logs the data to the console - a debugging operation
type ConsoleLogOperation struct{}

func (cl ConsoleLogOperation) Apply(ds *DataStream) *DataStream {
	output := make(chan []string)

	go func() {
		defer close(output)
		for data := range ds.data {
			fmt.Println(data)
			output <- data
		}
		fmt.Println("")
	}()

	return &DataStream{data: output}
}

func ConsoleLog() *ConsoleLogOperation {
	return &ConsoleLogOperation{}
}

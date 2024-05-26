package operations

// GetColumnsOperation gets a range of columns from the data stream - inclusive of the start index and exclusive of the end index
type GetColumnsOperation struct {
	startIndex int
	endIndex   int
}

func (gc *GetColumnsOperation) Apply(ds *DataStream) *DataStream {
	output := make(chan []string)
	go func() {
		defer close(output)
		for record := range ds.data {
			output <- record[gc.startIndex:gc.endIndex]
		}
	}()
	return &DataStream{data: output}
}

func GetColumns(startIndex int, endIndex int) *GetColumnsOperation {
	return &GetColumnsOperation{startIndex: startIndex, endIndex: endIndex}
}

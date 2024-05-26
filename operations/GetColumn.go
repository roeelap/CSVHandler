package operations

// GetColumnOperation gets a single column from the data stream
type GetColumnOperation struct {
	columnIndex int
}

func (gc *GetColumnOperation) Apply(ds *DataStream) *DataStream {
	output := make(chan []string)
	go func() {
		defer close(output)
		for record := range ds.data {
			output <- []string{record[gc.columnIndex]}
		}
	}()
	return &DataStream{data: output}
}

func GetColumn(columnIndex int) *GetColumnOperation {
	return &GetColumnOperation{columnIndex: columnIndex}
}

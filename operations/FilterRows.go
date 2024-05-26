package operations

// FilterRowsOperation filters rows based on a predicate
type FilterRowsOperation struct {
	predicate func([]string) bool
}

func (fr *FilterRowsOperation) Apply(ds *DataStream) *DataStream {
	output := make(chan []string)
	go func() {
		defer close(output)
		for record := range ds.data {
			if fr.predicate(record) {
				output <- record
			}
		}
	}()
	return &DataStream{data: output}
}

func FilterRows(predicate func([]string) bool) *FilterRowsOperation {
	return &FilterRowsOperation{predicate: predicate}
}

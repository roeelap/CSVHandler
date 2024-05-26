package operations

// DataStream represents a stream of data records that is passed through a pipeline of operations.
type DataStream struct {
	data chan []string
}

// Operation is an interface that represents an operation that can be applied to a DataStream.
// Every operation must implement the Apply method, which takes a DataStream as input and returns a DataStream as output.
// In order to create new operations, you need to define a struct that implements the Operation interface and provides an Apply method.
type Operation interface {
	Apply(ds *DataStream) *DataStream
}

// With applies an operation to a DataStream and returns the resulting DataStream.
func (ds *DataStream) With(operation Operation) *DataStream {
	return operation.Apply(ds)
}

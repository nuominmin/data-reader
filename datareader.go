package datareader

// DataProcessor 定义了解析数据和处理数据的接口
type DataProcessor[DATA, REQ, RESP any] interface {
	Parse(data []byte) (*DATA, error)
	HandleData(*DATA, *REQ) (*RESP, error)
}

// DataReader 是一个泛型接口，用于读取和解析数据
type DataReader[DATA, REQ, RESP any] interface {
	ReadData(*REQ) (*RESP, error)
}

// fileDataReader 实现了 DataReader 接口
type fileDataReader[DATA, REQ, RESP any] struct {
	data      *DATA
	processor DataProcessor[DATA, REQ, RESP]
}

// ReadData 调用传入的 processor 处理数据。
func (f *fileDataReader[DATA, REQ, RESP]) ReadData(r *REQ) (*RESP, error) {
	return f.processor.HandleData(f.data, r)
}

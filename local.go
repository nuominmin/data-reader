package datareader

import (
	"os"
)

// NewLocalReader 创建一个新的 fileDataReader 实例。
func NewLocalReader[DATA, REQ, RESP any](filePath string, processor DataProcessor[DATA, REQ, RESP]) (DataReader[DATA, REQ, RESP], error) {
	// 读取文件
	dataBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// 解析文件数据
	var d *DATA
	if d, err = processor.Parse(dataBytes); err != nil {
		return nil, err
	}

	return &fileDataReader[DATA, REQ, RESP]{
		data:      d,
		processor: processor,
	}, nil
}

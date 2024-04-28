package datareader

import (
	"encoding/json"
	"fmt"
	"testing"
)

// MyData 用于演示的数据结构
type MyData struct {
	data map[string]map[string]map[string]string // map[tx_hash] map[from] map[tick] amount
}

type Req struct {
	TxHash  string
	Address string
	Tick    string
}

type Resp struct {
	Address string
}

type jsonData struct {
	TxHash string `json:"tx_hash"`
	From   string `json:"from"`
	Amount string `json:"amount"`
	Tick   string `json:"tick"`
}

// JSONDataProcessor 实现了 DataProcessor 接口
type JSONDataProcessor struct{}

func (p *JSONDataProcessor) Parse(data []byte) (*MyData, error) {
	var result []jsonData
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	myData := &MyData{data: make(map[string]map[string]map[string]string)}

	var ok bool
	for i := 0; i < len(result); i++ {

		txHash := result[i].TxHash
		address := result[i].From
		tick := result[i].Tick

		if _, ok = myData.data[txHash]; !ok {
			myData.data[txHash] = make(map[string]map[string]string)
		}

		if _, ok = myData.data[txHash][address]; !ok {
			myData.data[txHash][address] = make(map[string]string)
		}

		if _, ok = myData.data[txHash][address][tick]; !ok {
			myData.data[txHash][address] = make(map[string]string)
		}

		myData.data[txHash][address][tick] = result[i].Amount
	}
	return myData, nil
}

func (p *JSONDataProcessor) HandleData(data *MyData, req *Req) (*Resp, error) {
	// 根据param逻辑定制数据返回，这里只是简单示例
	if _, ok := data.data[req.TxHash]; !ok {
		return nil, fmt.Errorf("no data for given param")
	}

	if _, ok := data.data[req.TxHash][req.Address]; !ok {
		return nil, fmt.Errorf("no data for given param")
	}

	if _, ok := data.data[req.TxHash][req.Address][req.Tick]; !ok {
		return nil, fmt.Errorf("no data for given param")
	}

	return &Resp{Address: data.data[req.TxHash][req.Address][req.Tick]}, nil
}

func TestDataRead(t *testing.T) {
	processor := &JSONDataProcessor{}

	dr, err := NewLocalReader[MyData, Req, Resp]("./data.json", processor)
	if err != nil {
		fmt.Println("Error initializing data reader:", err)
		return
	}
	resp, err := dr.ReadData(&Req{
		TxHash:  "0xfbffa89b0fab4ec5c21edf3ba8024f89985af750a61b6561c27e4918651ab8f7",
		Address: "0x4444777786851a1b941a86694f5f9a11da070f3f",
		Tick:    "Sparkle Inscription",
	}) // 使用"adult"作为param进行数据筛选
	if err != nil {
		t.Errorf("Error reading data: %v", err)
	}
	t.Logf("Read data: %+v\n", resp)
}

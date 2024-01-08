package strategy

import (
	"sync"
	"testing"
)

func TestStrategy(t *testing.T) {

	sf := StrategyFactory{
		mu: sync.Mutex{},
		m:map[string]IStrategy{
			"file":&FileStorage{},
			"encode":&EncodeStorage{},
		},

	}

	strategyType := "file"
	storage:=sf.NewStrategy(strategyType)
	storage.Store()

	strategyType2 := "encode"
	storage2:=sf.NewStrategy(strategyType2)
	storage2.Store()

}

// getData 获取数据的方法
// 返回数据，以及数据是否敏感
//func getData() ([]byte, bool) {
//	return []byte("test data"), false
//}

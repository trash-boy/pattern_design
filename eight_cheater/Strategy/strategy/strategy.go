package strategy

import (
	"fmt"
	"sync"
)

type IStrategy interface {
	Store()
}

type StrategyFactory struct {
	m  map[string]IStrategy
	mu sync.Mutex
}

type FileStorage struct {
}

func (f *FileStorage) Store() {
	fmt.Println("file文件存储")
}

type EncodeStorage struct {
}

func (e *EncodeStorage) Store() {
	fmt.Println("加密储存")
}

func (s *StrategyFactory) NewStrategy(key string) IStrategy {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.m[key]
}

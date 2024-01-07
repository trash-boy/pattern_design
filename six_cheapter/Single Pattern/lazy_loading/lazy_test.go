package lazy_loading

import (
	"sync"
	"sync/atomic"
	"testing"
)

//三种不同的lazy_loading

//1.读写锁
type Singleton struct {
	Name string
}

var (
	instance *Singleton
	lock sync.RWMutex
	flag uint32
	once sync.Once
)

func GetInstance1()*Singleton{
	lock.RLock()
	defer lock.RUnlock()

	if instance == nil{
		lock.Lock()
		defer lock.Unlock()
		if instance == nil{
			instance = &Singleton{"liuyang"}
		}
	}
	return instance
}

//2.利用原子操作
func GetInstance2()*Singleton{
	if atomic.LoadUint32(&flag) == 0{
		lock.Lock()
		defer lock.Unlock()
		if atomic.LoadUint32(&flag) == 0{
			instance = &Singleton{"liuyang1"}
			defer atomic.SwapUint32(&flag,1)
		}
	}
	return instance
}


//3.利用sync.Once
func GetInstance3()*Singleton {
	once.Do(func() {
		instance = &Singleton{
			"liuyang3",
		}
	})
	return instance
}

func BenchmarkV1Mutex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetInstance1()
	}
}

func BenchmarkV2Atomic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetInstance2()
	}
}

func BenchmarkV3SyncOnce(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetInstance3()
	}
}





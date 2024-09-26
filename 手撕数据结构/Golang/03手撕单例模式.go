//单例模式：确保一个类只有一个实例

package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	name string
}

var (
	instance *Singleton
	once     sync.Once
)

func (s *Singleton) GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {
	sing := &Singleton{}
	// 获取单例实例
	instance1 := sing.GetInstance()
	instance2 := sing.GetInstance()

	// 检查两个实例是否相同
	if instance1 == instance2 {
		fmt.Println("两个实例是相同的")
	} else {
		fmt.Println("两个实例是不同的")
	}
}

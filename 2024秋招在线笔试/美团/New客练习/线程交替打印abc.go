package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for {
		wg.Add(3)
		go task("a", &wg)
		time.Sleep(time.Second)
		go task("b", &wg)
		time.Sleep(time.Second)
		go task("c", &wg)
		time.Sleep(time.Second)
	}
}

func task(s string, wg *sync.WaitGroup) {
	fmt.Println(s)
	wg.Done()
}

package exercise

import (
	"fmt"
	"sync"
)

func MapConcurrencyTest() {
	testMap := make(map[string]uint32)
	var lock sync.Mutex
	for i := 0; i < 100; i++ {
		go func() {
			lock.Lock()
			testMap["num"]++
			lock.Unlock()
		}()
	}
	fmt.Println("test", testMap["num"])
}

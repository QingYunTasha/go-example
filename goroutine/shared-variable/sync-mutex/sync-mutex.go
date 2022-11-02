package syncMutex

import (
	"fmt"
	"sync"
	"time"
)

type safeNumber struct {
	v   int
	mux sync.Mutex
}

func Main() {
	total := safeNumber{v: 0}
	for i := 0; i < 1000; i++ {
		go func() {
			total.mux.Lock()
			total.v++
			total.mux.Unlock()
		}()
	}
	time.Sleep(time.Second)
	total.mux.Lock()
	fmt.Println(total.v)
	total.mux.Unlock()
}

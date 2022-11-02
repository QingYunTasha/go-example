package syncWaitGroup

import(
	"fmt"
	"time"
	"sync"
)


func Main(){
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go say("world", wg)
	go say("hello", wg)
	
	wg.Wait()
}

func say(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
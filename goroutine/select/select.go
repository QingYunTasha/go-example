package select_

import (
	"fmt"
	"time"
)

func Main() {
	ch := make(chan string)

	go func() {
		fmt.Println("calculate goroutine starts calculating")
		time.Sleep(time.Second) // Heavy calculation
		fmt.Println("calculate goroutine ends calculating")

		ch <- "FINISH"
		time.Sleep(time.Second)
		fmt.Println("calculate goroutine finished")
	}()

	for {
		select {
		case <-ch:
			fmt.Println("main goroutine finished")
			return
		default:
			fmt.Println("WAITING...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

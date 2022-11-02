package timeSleep

import (
	"fmt"
	"time"
)

func Main() {
	go say("world")
	go say("hello")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

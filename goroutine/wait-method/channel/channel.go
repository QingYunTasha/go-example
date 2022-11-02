package channel

import (
	"fmt"
	"time"
)

func Main(){
	ch := make(chan string)
	
	go say("world", ch)
	go say("hello", ch)
	
	<-ch
	<-ch
}


func say(s string, c chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	c <- "FINISH"
}
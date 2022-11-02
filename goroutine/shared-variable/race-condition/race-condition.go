package raceCondition

import (
	"time"
	"fmt"
)

func Main(){
	total := 0
	for i := 0; i < 1000; i++{
		go func(){
			total++
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(total)
}
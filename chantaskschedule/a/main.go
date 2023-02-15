package main

import (
	"fmt"
	"time"
)

func main() {
	chArr := [4]chan struct{}{
		make(chan struct{}),
		make(chan struct{}),
		make(chan struct{}),
		make(chan struct{}),
	}

	for i := 0; i < 4; i++ {
		go func(i int) {
			for {
				<-chArr[i%4]
				fmt.Printf("i am %d\n", i)

				time.Sleep(1 * time.Second)
				chArr[(i+1)%4] <- struct{}{}
			}
		}(i)
	}

	chArr[0] <- struct{}{}
	select {}
}

package main

import (
	"fmt"
	"time"
)

type Token struct{}

func newWorker(id int, ch chan Token, next chan Token) {
	for {
		token := <-ch
		fmt.Println(id + 1)
		time.Sleep(time.Second)
		next <- token
	}
}

func main() {
	chs := []chan Token{make(chan Token), make(chan Token), make(chan Token), make(chan Token)}

	for i := 0; i < 4; i++ {
		go newWorker(i, chs[i], chs[(i+1)%4])
	}

	chs[0] <- struct{}{}

	select {}
}

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var closing = make(chan struct{})
	var closed = make(chan struct{})

	go func() {
		// do biz
		for {
			select {
			case <-closing:
				return
			default:
				// biz working
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// 处理CTRL+C等中断信号
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	close(closing)
	// 执行退出之前的清理动作
	go doCleanup(closed)

	select {
	case <-closed:
	case <-time.After(time.Second):
		fmt.Println("clean up timeout, shutdown now")
	}
	fmt.Println("graceful shutdown")
}

func doCleanup(closed chan struct{}) {
	time.Sleep(time.Minute)
	close(closed)
}

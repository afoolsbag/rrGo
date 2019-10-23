package rrgo

import (
	"fmt"
	"testing"
	"time"
)

// 消息发送者
func messageSender(ch chan string) {
	ch <- "message 1" // 若通道已满，该协程将阻塞在此，直到通道不满
	ch <- "message 2"
	ch <- "message 3"
	ch <- "message 4"
	ch <- "message 5"
	close(ch) // 通道由生产者关闭
}

// 消息接收者
func messageReceiver(ch chan string) {
	// 若通道未关闭且通道已空，该协程将阻塞在此，直到通道不空
	for msg := range ch {
		fmt.Printf("Received message: %s\n", msg)
	}
}

// 协程
func TestCoroutines(t *testing.T) {
	ch := make(chan string) // 无缓冲通道

	go messageSender(ch)
	go messageReceiver(ch)
	ch <- "message 0"

	time.Sleep(1e9) // 休眠等待 1s 使协程有机会执行
}

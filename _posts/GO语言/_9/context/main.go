package main

import (
	"context"
	"fmt"
	"time"
)

func HandelRequest(ctx context.Context) {
	go WriteLog(ctx)
	go WriteDB(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("请求处理完毕")
			return
		default:
			fmt.Println("请求处理中……")
			time.Sleep(2 * time.Second)
		}
	}
}
func WriteLog(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("写日志完成")
			return
		default:
			fmt.Println("写日志中……")
			time.Sleep(2 * time.Second)
		}
	}
}
func WriteDB(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("写数据库完成")
			return
		default:
			fmt.Println("写数据库中……")
			time.Sleep(2 * time.Second)
		}
	}
}
func main() {
	//WithCancel
	/* ctx, cancel := context.WithCancel(context.Background())
	go HandelRequest(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("所有子协程都需要结束!")
	cancel()
	//Just for test whether sub goroutines exit or not
	time.Sleep(5 * time.Second) */

	//WithTimeout
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	go HandelRequest(ctx)
	time.Sleep(10 * time.Second)

}

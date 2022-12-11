package main

import (
	"context"
	"fmt"
	"time"
)

func HandelRequest(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("请求处理完毕")
			return
		default:
			fmt.Println("请求处理中……, parameter: ", ctx.Value("parameter"))

			time.Sleep(2 * time.Second)
		}
	}
}
func main() {
	ctx := context.Background()
	go HandelRequest(ctx)

	time.Sleep(2 * time.Second)
	ctx1 := context.WithValue(ctx, "parameter", "1")
	go HandelRequest(ctx1)
	time.Sleep(4 * time.Second)

}

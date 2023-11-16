package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	var wg sync.WaitGroup
	wg.Add(2)
	// timeout の時間によって ch1 , ch2 の読み込みが終わる前にタイムアウトをさせたりできる
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// それぞれの channel に値を送信する goroutine を起動
	go func() {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)
		ch1 <- "あああ"
	}()
	go func() {
		defer wg.Done()
		time.Sleep(800 * time.Millisecond)
		ch2 <- "いいい"
	}()

	// for 文に loop というラベルを付けて、break loop で for 文を抜ける
loop:
	for ch1 != nil || ch2 != nil {
		select {
		case <-ctx.Done():
			fmt.Println("キャンセルされました")
			break loop
		case <-ch1:
			println("ch1 から受信")
			ch1 = nil
		case <-ch2:
			println("ch2 から受信")
			ch2 = nil
		}
	}

	wg.Wait()
	fmt.Println("終了")

}

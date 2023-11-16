package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("あああ")
	}()

	// wait がないと、子のgoroutine が終了する前に main 関数が終了してしまう
	wg.Wait()
	fmt.Println("いいい")
}

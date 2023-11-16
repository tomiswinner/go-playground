package main

import (
	"fmt"
	"sync"
)

func main() {
	for true {
		var i int
		var wg sync.WaitGroup
		var mu sync.Mutex
		wg.Add(16)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			i++
		}()
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			i++
		}()
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			i++
		}()
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			i++
		}()
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			i++
		}()
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			i++
		}()
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			i++
		}()
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			i++
		}()
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			i++
		}()
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			i++
		}()
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			i++
		}()
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			i++
		}()
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			i++
		}()
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			i++
		}()
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			i++
		}()
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			i++
		}()

		wg.Wait()
		// i は、レースコンディションを、mutex で排他制御しているので、必ず 16 になる。
		fmt.Println(i)
	}
}

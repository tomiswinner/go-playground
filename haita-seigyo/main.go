package main

import (
	"fmt"
	"sync"
)

func main() {
	for true {
		var i int
		var wg sync.WaitGroup
		wg.Add(16)
		go func() {
			defer wg.Done()
			i++
		}()
		go func() {
			defer wg.Done()
			i++
		}()
		go func() {
			defer wg.Done()
			i++
		}()
		go func() {
			defer wg.Done()
			i++
		}()
		go func() {
			defer wg.Done()
			i++
		}()
		go func() {
			defer wg.Done()
			i++
		}()
		go func() {
			defer wg.Done()
			i++
		}()
		go func() {
			defer wg.Done()
			i++
		}()
		go func() {
			defer wg.Done()
			i++
		}()
		go func() {
			defer wg.Done()
			i++
		}()
		go func() {
			defer wg.Done()
			i++
		}()
		go func() {
			defer wg.Done()
			i++
		}()
		go func() {
			defer wg.Done()
			i++
		}()
		go func() {
			defer wg.Done()
			i++
		}()
		go func() {
			defer wg.Done()
			i++
		}()
		go func() {
			defer wg.Done()
			i++
		}()

		wg.Wait()
		// i は、レースコンディションにより、0〜16の間の値が出力される
		fmt.Println(i)
	}
}

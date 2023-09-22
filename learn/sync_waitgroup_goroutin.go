package main

import (
	"fmt"
	"sync"
)

func printText(text string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
	}

	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)
	go printText("Hello", &wg)
	go printText("Selamat datang", &wg)

	wg.Wait()

}

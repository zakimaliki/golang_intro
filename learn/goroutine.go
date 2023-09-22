package main

// import (
// 	"fmt"
// 	"time"
// )

// func prinrSalam(text string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(text)
// 	}
// }

// func prinrText() {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(500 * time.Millisecond)
// 		fmt.Println("text", i)
// 	}
// }

// func prinrAngka() {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(200 * time.Millisecond)
// 		fmt.Println(i)
// 	}
// }

// func main() {
// 	// runtime.GOMAXPROCS(4)// jumlah core

// 	start := time.Now()

// 	// go prinrSalam("Hello")
// 	// prinrSalam("Selamat datang")

// 	go prinrAngka()
// 	go prinrText()
// 	time.Sleep(100 * time.Millisecond)

// 	fmt.Println(time.Since(start))

// }

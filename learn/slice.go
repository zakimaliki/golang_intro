package main

// import (
// 	"fmt"
// )

// func main() {
// 	var fruits = []string{"apple", "grape", "banana", "melon"}
// 	var newFruits = fruits[0:2]
// Kode	Output	Penjelasan
// fruits[0:2]	[apple, grape]	semua elemen mulai indeks ke-0, hingga sebelum indeks ke-2
// fruits[0:4]	[apple, grape, banana, melon]	semua elemen mulai indeks ke-0, hingga sebelum indeks ke-4
// fruits[0:0]	[]	menghasilkan slice kosong, karena tidak ada elemen sebelum indeks ke-0
// fruits[4:4]	[]	menghasilkan slice kosong, karena tidak ada elemen yang dimulai dari indeks ke-4
// fruits[4:0]	[]	error, pada penulisan fruits[a:b] nilai a harus lebih kecil atau sama dengan b
// fruits[:]	[apple, grape, banana, melon]	semua elemen
// fruits[2:]	[banana, melon]	semua elemen mulai indeks ke-2
// fruits[:2]	[apple, grape]	semua elemen hingga sebelum indeks ke-2

// fmt.Println(newFruits)
// fmt.Println(len(fruits))

// fmt.Println(len(fruits)) // len: 4
// fmt.Println(cap(fruits)) // cap: 4

// var aFruits = fruits[0:3]
// fmt.Println(len(aFruits)) // len: 3
// fmt.Println(cap(aFruits)) // cap: 4

// var bFruits = fruits[1:4]
// fmt.Println(len(bFruits)) // len: 3
// fmt.Println(cap(bFruits)) // cap: 3

// 	Kode	Output	len()	cap()
// fruits[0:4]	[buah buah buah buah]	4	4
// aFruits[0:3]	[buah buah buah ----]	3	4
// bFruits[1:4]	---- [buah buah buah]	3	3

// var cFruits = append(fruits, "papaya")
// fmt.Println(cFruits)

// dst := make([]string, 3)
// src := []string{"watermelon", "pinnaple", "apple", "orange"}
// n := copy(dst, src)
// fmt.Println(dst) // watermelon pinnaple apple
// fmt.Println(src) // watermelon pinnaple apple orange
// fmt.Println(n)   // 3

// 3 index adalah teknik slicing elemen yang sekaligus menentukan kapasitasnya.
// Cara menggunakannnya yaitu dengan menyisipkan angka kapasitas di belakang,
// seperti fruits[0:1:1]

// var fruits = []string{"apple", "grape", "banana"}
// var aFruits = fruits[0:2]
// var bFruits = fruits[0:2:2]

// fmt.Println(fruits)      // ["apple", "grape", "banana"]
// fmt.Println(len(fruits)) // len: 3
// fmt.Println(cap(fruits)) // cap: 3

// fmt.Println(aFruits)      // ["apple", "grape"]
// fmt.Println(len(aFruits)) // len: 2
// fmt.Println(cap(aFruits)) // cap: 3

// fmt.Println(bFruits)      // ["apple", "grape"]
// fmt.Println(len(bFruits)) // len: 2
// fmt.Println(cap(bFruits)) // cap: 2
// }

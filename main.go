// Package Name
package main

// Import
import "fmt"

func main() {

	// This is Commentary

	// Print New Line
	fmt.Println("Hello, World!")

	// Print In Line
	fmt.Print("Hello, ")
	fmt.Println("world!")

	// Declared Variabel
	var firstName string = "john"

	// Declared and Define Variabel
	var lastName string
	lastName = "wick"

	// Print Format
	fmt.Printf("halo %s %s!\n", firstName, lastName)

	// Print
	fmt.Printf("halo john wick!\n")
	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName+"!")
	fmt.Println("halo", firstName, lastName, "!")

	// Infered Variable
	var firstName2 = "something"
	firstName3 := "someting_2"

	fmt.Println(firstName2)
	fmt.Println(firstName3)

	// Multi Variable
	// var first, second, third string
	// first, second, third = "satu", "dua", "tiga"

	// var fourth, fifth, sixth string = "empat", "lima", "enam"

	// seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	// one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	// Reserve Varible
	// _ can't be get or printed, used when there is return in function but we don't the value
	_ = "something"
	a, _ := "abc", "DEF"
	fmt.Print(a)

	// Pointer Variable
	// name := new(string)

	// fmt.Println(name) // 0x20818a220
	// fmt.Println(*name)

	// make()
	// channel
	// slice
	// map

	// const
	//const firstName string = "john"
	//fmt.Print("halo ", firstName, "!\n")

	// infered
	// const lastName = "wick"
	// fmt.Print("nice to meet you ", lastName, "!\n")

	// const (
	// 	square         = "kotak"
	// 	isToday  bool  = true
	// 	numeric  uint8 = 1
	// 	floatNum       = 2.2
	// )

	// const (
	// 	a = "konstanta"
	// 	b
	// )

	// const (
	// 	today string = "senin"
	// 	sekarang
	// 	isToday2 = true
	// )

	// const satu, dua = 1, 2
	// const three, four string = "tiga", "empat"
}

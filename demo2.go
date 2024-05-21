package main

import (
	"fmt"
	// "io"
	// "math"
	// "os"
)

func main() {
	// var dd int = 17
	// var mm int = 04
	// var yy int = 2023
	// var str string
	// str = fmt.Sprintf("%02d-%02d-%04d", dd, mm, yy) //Sprintf to define the date it give the output in the form of string
	// io.WriteString(os.Stdout, str)                  // WriteString function is to print the output in the desired location
	// fmt.Print("\n")
	// var str1 string
	// fmt.Printf("Enter the String :")
	// fmt.Scan(&str1) // Dont need a format specifier
	// fmt.Printf("Result: %s\n", str1)
	// var str2 string
	// fmt.Printf("Enter the Name :")
	// fmt.Scan(&str2) // Dont need a format specifier
	// fmt.Printf("Result: %s\n", str2)
	// var num1 int
	// var num2 float64
	// fmt.Scan(&num1, &num2)                             // Dont need a format specifier
	// fmt.Printf("Result: %d %f %s\n", num1, num2, str2) //slide number 120
	// fmt.Println(str)
	// var a uint = 123
	// var b int = 0
	// b = int(a)
	// fmt.Println(a, b)
	// var x int = 255
	// var r float32
	// r = float32(math.Sqrt(float64(x)))
	// fmt.Println(r)
	var x int = 42
	var y float64 = float64(x)
	var z uint = uint(y)
	fmt.Printf("Value of X is %d and type is %T\n", x, x)
	fmt.Printf("Value of Y is %.2f and type is %T\n", y, y)
	fmt.Printf("Value of Z is %d and type is %T\n", z, z)

}

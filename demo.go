package main

import (
	"fmt"
)

var chandra = 20 // Global Funtion var Local function :=
func main() {
	// after removing the int it will be inferred
	// var a, b, c, d = 1, 3.5, 5.4, "chandra" // declaring the multiple variables
	// e, f := 7, "Hello Mahima"

	// Multiple variable declaration can also be grouped together into a block for greater readabilty
	// var (
	// 	g int
	// 	h int  = 1
	// 	i string = "hello"
	// )
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)
	// fmt.Println(f)
	// fmt.Println(g)
	// fmt.Println(h)
	// fmt.Println(i)
	// fmt.Printf("%d\n",a) // %d is format specifier
	// fmt.Printf("%d\n",h)
	// fmt.Println(a,b,c)
	// fmt.Printf("%d,%f,%d,%d",a,b,e,h)
	// fmt.Println(chandra)

	fmt.Println("Hello")

	// Scanf for first string
	// var name string = "A Chandr Kishan Singha"
	// var name1 string;
	// var age int = 20
	// var branch string = "Computer Science and Engineering (AIML)"
	// var college string = "Shri Ramdeobaba College of Engineering and Management"
	// fmt.Println("Enter Your Name : ")
	// fmt.Scanf("%s %s",&name,&name1)
	// fmt.Println("Hello ", name ,name1)
	// fmt.Print("Enter First String :")
	// fmt.Scanln(&name)
	// fmt.Print("Enter Second String :")
	// fmt.Scanln(&name1)
	// fmt.Println(name," "+ name1)
	// var a,b int=10,20
	// fmt.Printf("%d %d\n",a,b)
	// fmt.Printf("%d",a+b)
	// Go has three funtion to outpt text
	// print- No new line is generate and have no format specifier
	// printf- Have format specifier
	// println- create new line and does not have format specifier
	// fmt.Println("Name : ", name)
	// fmt.Println("Age : ", age)
	// fmt.Println("Branch : ", branch)
	// fmt.Println("College Nme : ", college)
	// fmt.Print("\n")
	// fmt.Println("Name : ", name, "                        ", "Age: ", age)
	// fmt.Println("Branch : ", branch, "      College Name: ", college)

	// var a string = "Rajesh"
	// fmt.Printf("Value of a is '%s'", a)

	// const (
	// 	name = "Geeks for Geeks"
	// 	dept = "CS"
	// )

	// fmt.Printf("%s is a %s Portal.\n", name, dept)
	// var num1 int = 21
	// var num2 float32 = 7.798
	// var num3 int = 14
	// var num4 = 4 + 1i
	// fmt.Printf("%d\n", num1) // Integer
	// fmt.Printf("%g\n", num2) // Floating
	// fmt.Printf("%b\n", num3) // Binray
	// fmt.Printf("%e\n", num4) // Imaginary Number
	// fmt.Printf("%s\n", name) // String
	// fmt.Printf("%v\n%v\n%v", num1, num2, name)

	// Lvalue in Rvalue
	// Integer value literal
	// decimal 10
	// Binary 2   starts with "0b" or "0B"  eg:
	// Octal 8 starts with "0" or "0o" or "0O"
	// HexaDecimal  start with "0x" and "0X"

	// fmt.Println(15 == 017)
	// fmt.Println(15 == 0xF)

	var a int = 10
	// var b int =20
	// a++
	// fmt.Printf("%d",a)
	// fmt.Println(a--)
	// fmt.Println(--a)
	// fmt.Println(++a)

	fmt.Printf("Binary of %d is %0b\n", a, a)
	fmt.Printf("Heaxdecimal of %d is %0x\n", a, a)
	fmt.Printf("Octal of %d is %0o\n", a, a)

	// format specifier
	// 	v = for all
	// 	d = integer
	//  b = binary
	// 	o = octal
	//  g = float
	//  t = boolean
	//  s = string
	//  T = type

	fmt.Printf("Expresion is 15==0xf is %T\n", 15 == 0xf)

	var i, j string = "Hello", "World"
	fmt.Printf("%s\n", i)
	fmt.Printf("%s\n", j)
	fmt.Print(i, "\n")
	fmt.Print(j, "\n")
	fmt.Printf("%s %s\n", i, j) // inserts a space between the arguments if neither are strings
	fmt.Print(20, "Rajesh")
	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%b\n", a)
	fmt.Printf("%d\n", a)
	fmt.Printf("%+d\n", a) // + is the sign
	fmt.Printf("%o\n", a)
	fmt.Printf("%x\n", i)
	fmt.Printf("%X\n", i)
	fmt.Printf("%#x\n", i)
	fmt.Printf("%4d\n", a)  // digit alloation
	fmt.Printf("%04d\n", a) // digit allocation and rest of the space are filled by 0

}

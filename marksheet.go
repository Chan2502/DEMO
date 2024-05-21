package main

import (
	"fmt"
)

func main() {
	var name string
	var rollno int
	var math int
	var science int
	var english int
	var hindi int
	var computer int
	var total int
	var percentage float64

	fmt.Printf("Enter the Name :")
	fmt.Scan(&name)
	fmt.Printf("Enter the Roll No :")
	fmt.Scan(&rollno)
	fmt.Printf("Enter the Marks of Math :")
	fmt.Scan(&math)
	fmt.Printf("Enter the Marks of Science :")
	fmt.Scan(&science)
	fmt.Printf("Enter the Marks of English :")
	fmt.Scan(&english)
	fmt.Printf("Enter the Marks of Hindi :")
	fmt.Scan(&hindi)
	fmt.Printf("Enter the Marks of Computer :")
	fmt.Scan(&computer)
	total = math + science + english + hindi + computer
	percentage = float64(total) / 5
	fmt.Print("\n")
	fmt.Println("------------------------------------------------MARKSHEET------------------------------------------------")
	fmt.Print("\n")
	fmt.Printf("Roll No : %d                                                                              ", rollno)
	fmt.Printf("Name : %s\n", name)
	fmt.Print("\n")
	fmt.Println("---------------------------------------------------------------------------------------------------------")
	fmt.Print("\n")
	fmt.Print("|      Course Code       |          Course Name         |           Marks          |        Grade        |\n")
	fmt.Println("---------------------------------------------------------------------------------------------------------")
	fmt.Print("\n")
	fmt.Printf("|        MAT102          |             MATHS            |            %d            |         AA         |\n", math)
	fmt.Print("\n")
	fmt.Printf("|        MAT102          |            SCIENCE           |            %d            |         AA         |\n", science)
	fmt.Print("\n")
	fmt.Printf("|        MAT102          |            ENGLISH           |            %d            |         AB         |\n", english)
	fmt.Print("\n")
	fmt.Printf("|        MAT102          |             HINDI            |            %d            |         AA         |\n", hindi)
	fmt.Print("\n")
	fmt.Printf("|        MAT102          |            COMPUTER          |            %d            |         AA         |\n", computer)
	fmt.Print("\n")
	fmt.Println("---------------------------------------------------------------------------------------------------------")
	fmt.Print("\n")
	fmt.Printf("Total Marks : %d                                                                    ", total)
	fmt.Printf("Percentage : %.2f %%\n", percentage)

}

package main

import "fmt"

func main() {
	var first_number int
	var second_number int
	fmt.Print("Enter a large number: ")
	fmt.Scan(&first_number)
	fmt.Print("Enter a smaller number: ")
	fmt.Scan(&second_number)
	fmt.Println("Remainder of ", first_number, "/", second_number, " is: ", first_number%second_number)
}
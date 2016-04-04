package main

import "fmt"

func main() {
    my_int, my_bool, my_string := 9, true, "Testing... 1, 2, 3."
    
	fmt.Printf("%T \n", my_int)
	fmt.Printf("%T \n", my_bool)
	fmt.Printf("%T \n", my_string)
}

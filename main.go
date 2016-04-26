package main

import (
	"fmt"
	"strconv"
)

func main() {

	A := 43

	var B = &A // 0xc082002278 - Memory Address

	stringedAddressA := fmt.Sprintf("%p", &A) // Converted to string for print purposes

	fmt.Println(strconv.Itoa(A) + " - Value of A")           // Prints 43
	fmt.Println(stringedAddressA + " - Memory Address of A") // Prints 0xc082002278

	*B = 90 // Changing the VALUE stored at memory address 0xc082002278

	stringedAddressB := fmt.Sprintf("%p", B)

	fmt.Println("------------------------------------------------")                      // <-Just a 'Before & After' separator
	fmt.Println(stringedAddressB + " - Memory Location of A via B Pointer")              // Prints 0xc082002278
	fmt.Println(strconv.Itoa(*B) + " - Changed Value Now Stored At Memory Address of A") // Prints 90
	fmt.Println(strconv.Itoa(A) + " - Value of A")                                       // Prints 90 [original value 43 is gone, we swapped 43 out from underneath A and put 90 in its place]
	fmt.Println(stringedAddressA + " - Memory Address of A")                             // Prints 0xc082002278, same memory address

	// the above code makes B a pointer to the memory address where an int is stored
	// B is of type "int pointer"
	// *int -- the * is part of the type -- B is of type *int
}

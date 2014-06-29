/*
Problem #34: Stimulating Call By Reference
Author: Corey Prak
Date Created: 06/03/2014

Desecription: Write a program that contains a main function and a function called swap. Your main program should read in two integer values from the keyboard and then call the function swap.  The swap function will swap the original values and then return back to main. The function main wil then output the values contained in the original variables. The values should be swapped. You must stimulate call by reference by passing the address of the variables into the function swap.

Comments:

- Definitely needs an extensive writeup for newbies

- Note how the package was imported, and how the content of the fmt package was accessed.
*/

package main

import . "fmt"

func swap(num1 *int, num2 *int) {
	var numberHolder int

	numberHolder = *num1

	*num1 = *num2
	*num2 = numberHolder
}

func main() {
	var numberInput1, numberInput2 int

	Printf("\nPlease enter two integer values:")

	Scanf("%d %d", &numberInput1, &numberInput2)

	Printf("\n\nThe values of the two integers received: ")
	Printf("\nnumberInput1 = %d", numberInput1)
	Printf("\nnumberInput2 = %d", numberInput2)

	swap(&numberInput1, &numberInput2)

	Printf("\n\nThe values of the two integers after calling the swap function: ")
	Printf("\nnumberInput1 = %d", numberInput1)
	Printf("\nnumberInput2 = %d\n\n", numberInput2)
}

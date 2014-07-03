/*******************************************************************************
Problem #37: Perfect Numbers
Author: Corey Prak
Date Created: 07/03/2014

 Description: Greek Mathematicians took a special interest in numbers that are equal to the sum of their proper divisors. A proper divisor of n is any number that evenly divides n that is less than n itself. For example, 6 is a perfect number because it is the sum of 1, 2, and 3 which are the factors of 6 that are less than 6. In other words, when you add up the factors of a number, excluding the number itself, you get the number.
Similarly, the number 28 is a perfect number because 1 + 2 + 4 + 7 + 14 add up to 28.
 You are to write a program that contains a function called IsPerfect that takes an integer n and returns true (1) if n is perfect and false (0) otherwise. Your program should output all the perfect numbers between 1 and 100,000.

Comments:
 - Warmed up, wanted to time it for fun! Started 8:25AM. Finished 8:54AM

- Most of the problem came from attempting to handle the divisor receiving 
 a value of 0. I initialized the condition in divisors to be < instead of <=.
This removed 1, and so I went down a rabbit hole of ensuring 1 was a perfect 
number. Then I realized, NO, it isn't. 1 is the dang number, it must be 
excluded! After realizing that, I understood that I somehow needed to handle 0
not being a perfect number. If we call IsPerfect on 0, the Divisors function 
would return a nil slice, which means that nothing would be iterated on at 
IsPerfect, so sum would stay at its default value of 0. Because IsPerfect 
checks to see whether the sum of a number n's divisors is equal to it, the
function would return true on 0. I added in a condition at the end to 
handle that. THEN, I realized that the assignment specified that we needed
the for loop in main to start at 1... As you can see, the biggest obstacle in 
not completing this earlier was because I didn't really pay attention to the
requirements. A bit embarassing, but I'll try to be more aware of when I do
this. Hopefully I'll catch myself. 
*******************************************************************************/

package main

import "fmt"

// Go WILL maintain memory that may be out of the scope of the block 
// it was declared in, so long as something is referencing it.
func divisors(n int) []int {
	var intSlice []int

		for i := 1; i < n; i++ {
		if (n % i) == 0 {
			intSlice = append(intSlice, i)
		}
	}

	return intSlice
}

func IsPerfect(n int) bool {
	var sum int
	
		for _, value := range divisors(n) {
		sum += value
	}

	//Takes care of 0, but this will never be true since i in main() is 
	//starting at 1. Refer to comments on why I left it here.
	return (sum == n) && (n != 0)
}

func main() {
	fmt.Printf("\n")

	for i := 0; i <= 100000; i++ {
		if IsPerfect(i) {
			fmt.Printf("%d  ", i)
		} 
	}

	fmt.Printf("\n\n")
}
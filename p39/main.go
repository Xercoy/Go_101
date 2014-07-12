/*******************************************************************************
Problem #39: Left Shift Plus
Author: Corey Prak
Date Created: ~06/08/2014
Date Completed: 06/11/2014
Comments:
- No need for %u format specifier! If there is an type that is unsigned, it will
be printed as such. Awesome!

- Placing number after % indicates width
- Placing a . and a number after % indicates precision. 
- The two can be combined, ex. %9.2f (width 9, precision 2)
- Measured in Unicode points, and not numbers of bytes like Printf in C
*******************************************************************************/

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x int
	var y uint
	var limit uint

	x, y = 1, 1
	limit = uint(unsafe.Sizeof(x) * 8)

		fmt.Printf("\nThe variable limit = %d.\n", limit)

	fmt.Printf("\nInitial values of x=%d and y=%d\n", x, y)

	fmt.Printf("\n%12c %12c\n", 'X', 'Y')

		for i := 0; i <= int(limit-1); i++ {
		x = x << 1
		y = y << 1
		fmt.Printf("%12d %12d\n", x, y)
	}
}

/*Writeup:

 The variables x(int) and y(uint) are assigned a default value of 1, while the size of an integer is 4 bytes, meaning the variable limit is equal to 32. 
 
 The for loop is shifting the binary value of 1 and the remaining bits to the left, increasing the value by powers of two at each iteration of the for loop. 

 Before the values of x and y are 0, except for x being negative, they are the same integer value of 2^31. This is because the signed bit that is on the leftmost bit for x indicates that the integer should be negative. Since y is an unsigned integer, the leftmost bit does not represent sign. 

 Finally, both values are zero at the final iteration because the only bit that is "on" is shifted to the left and truncated. Every bit is 0.
 
*/
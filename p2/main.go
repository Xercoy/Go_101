/*
Problem #2: The Value 6
Author: Corey Prak
Date Created: 04/03/2014
Comments:
06/02/2014
The function had an error, I had used Println instead of Printf, resulting in incorrect output.

Added descriptive comments.

Note the way that the fmt package was imported, and how its content was accessed. We gave the "fmt" a nickname.

Additional notes: If %d was not given an expression of a matching type the program would not crash, but the Printf functino would indeed complain. Being the nice function that it is, it will replace the format specifier with an error message. An expression of type 6.2 will yield the error message: %!d(float64=6.2)

! is somewhat synonymous with the word no. In this case, we see that %!d can be inferred to mean, "Not an expression of type integer"

Being an even nicer function, it provides helpful information to us to convey what the problem is. (float64=6.2) tells us that the number 6.2 is of type float64.
*/

package main

import foo "fmt"

func main() {

	/* Similar to Problem 1, however the string literal includes a %d which is a special sequence of characters called a format specifier that the Printf understands. Whenever Printf receives a format specifier, it also expects to receive an additional argument which will replace (and sometimes be formatted by) the format specifier. Format specifers have a particular condition in which the expression of the argument must match a particular type. In our case, %d expects an expression to be an integer type, which is exactly what we did by making 6 the next argument. */
	foo.Printf("%d", 6)
}

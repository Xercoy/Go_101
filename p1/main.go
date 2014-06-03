/*
Problem #1: Hello World
Author: Corey Prak
Date Created: 04/03/2014
Comments:
06/02/2014 - Added comments to describe the program. 
*/

/* Every program must include a main package, along with a function named main 
   that does not take any arguments. */
package main

/* Import the 'fmt' package. All packages must be included with quotes around it */
import "fmt"

func main() {

/* We are calling the Println that is apart of the fmt package. As you can see, 
   the name of the package and a period comes before the function. There are more 
   ways to access content of a package, more on that later. Lastly, we provide the 
   string literal "Hello, world!" */
   fmt.Println("Hello, world!")
}

/*
Problem #3: The Character P
Author: Corey Prak
Date Created: 04/03/2014
Comments:

06/02/2014
Added comments to describe code, made a few changes to the code. 

Note yet another way to refer to the content of a package. For the other two, refer to Problems 1 and 2. 

Concerning data types and interpretation, there is much more to the topic than I had thought which goes beyond the scope of this program. I will work on a separate article with additional examples. 

A byte is an unsigned 8 bit integer. 
*/

package main

import . "fmt"

func main(){

  // var character byte = 80 will also work, note that a byte is an alias to uint8. Don't stress if you don't understand why just yet. 
  var character byte = 'P'

  /* Much like Problem 1 and 2, except this time we're using a different format specifier. */
  Printf("%c", character)
}

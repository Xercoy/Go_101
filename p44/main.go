/*******************************************************************************
Problem #43: Random Numbers 2
Author: Corey Prak
Date Created: 04/03/2014
Comments: I would have liked to call myRand() ten times from main instead, but that's not possible since there's not a convenient way to generate a new Source without keeping this exercise simple.

Description: Define a function myRand which returns random number between 19 and -19. The program should output 10 numbers.
*******************************************************************************/
package main

import (
	"fmt"
	"math/rand" // http://golang.org/pkg/math/rand/
	"time"
)

func myRand() {

	/* Specify the seed to the default Source, if it isn't called then he Source will seed with 1. Here we used time.Now so that the seed is different every time the program is run. t.Unix returns number of milaseconds since 1/1/1970*/
	rand.Seed(time.Now().Unix())

	/* The argument to Intn specifies the upper limit of numbers that can be generated. We subtract 19 from the number returned, effectively changing the range from 0 - 39 to -19 - 19*/
	for i := 0; i < 9; i++ {
		fmt.Println(rand.Intn(39) - 19)
	}
}

func main() {
	myRand()
}

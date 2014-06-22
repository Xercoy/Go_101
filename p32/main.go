/*******************************************************************************
 Problem #32: Inner Product of Two Vectors
Author: Corey Prak
Date Created: 06/22/2014
Comments:
first file of many being used with gofmt.
I tried my best to DRY the code while keeping it simple and readable.
*******************************************************************************/
package main

import (
	"fmt"
	"log"
)

func inner(u []float32, v []float32, size int) float32 {
	var innerProduct float32

	for k := 0; k <= size; k++ {
		innerProduct += u[k] * v[k]
	}

	return innerProduct
}

func main() {
	var firstVector, secondVector []float32
	var integerFromUser float32

	firstVector = make([]float32, 8)
	secondVector = make([]float32, 8)

	for i := 1; i <= 2; i++ {
		fmt.Printf("\nPlease enter 8 float32ing point numbers for the")

		if i == 1 {
			fmt.Printf(" first ")
		} else {
			fmt.Printf(" second ")
		}

		fmt.Printf("vector.\n\n")

		for j := 0; j <= 7; j++ {
			_, scanError := fmt.Scanf("%f", &integerFromUser)
			if scanError != nil {
				log.Fatalln(scanError)
			}

			if i == 1 {
				firstVector[j] = integerFromUser
			} else {
				secondVector[j] = integerFromUser
			}
		}
	}

	fmt.Printf("\nThe inner product of the \nfirst vector:%v", firstVector)
	fmt.Printf(" and the \nsecond vector %v \nis %f.\n\n",
		secondVector, inner(firstVector, secondVector, len(secondVector)-1))
}

/*******************************************************************************
 Problem #31: Bubble Sort
Author: Corey Prak
Date Created: 06/20/2014
Comments:
*******************************************************************************/

package main

import (
	"fmt"
	"log"
	"os"
)

func bubbleSort(slice []int) {
	for end := (len(slice) - 1); end >= 0; end-- {
		for iteratingSlice := slice; &iteratingSlice[0] != &slice[end]; iteratingSlice = iteratingSlice[1:] {
			if iteratingSlice[0] > iteratingSlice[1] {
				elementHolder := iteratingSlice[0]
				iteratingSlice[0] = iteratingSlice[1]
				iteratingSlice[1] = elementHolder
			}
		}
	}
}

func main() {
	numberSlice := make([]int, 10)
	n := 0

	file, err := os.Open("testdata31")
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i <= 9; i++ {
		fmt.Fscanf(file, "%d", &n)
		numberSlice[i] = n
	}

	fmt.Printf("\nSlice read from file before sort: %v\n", numberSlice)

	bubbleSort(numberSlice)

	fmt.Printf("\nSlice after bubble sort: %v\n\n", numberSlice)

	file.Close()
}

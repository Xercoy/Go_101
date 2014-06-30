/*******************************************************************************
Problem #35: Passing a Two Dimensional Array
Author: Corey Prak
Date Created: 06/29/2014
Comments:
*******************************************************************************/
package main

import (
	"fmt"
	"log"
)

func sum(array [3][3]int) int {
	var i, j, sum int

	for ; i < 3; i++ {
		for ; j < 3; j++ {
			sum += array[i][j]
		}
		j = 0
	}

	return sum
}

func main() {
	var tDArray [3][3]int
	var i, j int

	for ; i < 3; i++ {
		for ; j < 3; j++ {
			fmt.Printf("\nPlease enter an integer value for the 3x3 array in position [%d][%d]: ", i, j)
			_, scanErr := fmt.Scanf("%d", &tDArray[i][j])
			if scanErr != nil {
				log.Fatalln(scanErr)
			}
		}
		j = 0
	}

	fmt.Println("\nThe complete array: ", tDArray)

	fmt.Printf("\nThe sum of all the elemtns is %d.\n\n", sum(tDArray))
}

/*******************************************************************************
 Problem #38: Blank Removal
Author: Corey Prak
Date Created: 07/06/2014
Comments:

32 is ASCII equiv. of blank
*******************************************************************************/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	previousCharacterWasBlank := false

	fileContents, readErr := ioutil.ReadFile("testdata38")
	if readErr != nil {
		log.Fatalln(readErr)
	}

	for i := 0; i < len(fileContents); i++ {
		if fileContents[i] == 32 {

			if previousCharacterWasBlank {
				continue
			} else {
				fmt.Printf("%c", fileContents[i])
				previousCharacterWasBlank = true
			}

		} else {

			fmt.Printf("%c", fileContents[i])

			previousCharacterWasBlank = false

		}
	}

}

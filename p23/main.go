/*
Problem #23: Scanf Returns What?
Author: Corey Prak
Date Created: 05/10/2014
Comments:

func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
Fscanf scans text read from r, storing successive space-separated values into successive arguments as determined by the format. It returns the number of items successfully parsed.

*/

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var reader io.ReadCloser
	var openErr, closeErr, scanErr error
	var data int
	var successfulParses, totalSuccessfulParses int

	reader, openErr = os.Open("testdata23")
	if openErr != nil {
		fmt.Printf("\n\nOpen Error: %v\n\nExiting...\n\n")
		return
	}

	fmt.Printf("\nContents of file testdata23:\n")

	for scanErr == nil {
		successfulParses, scanErr = fmt.Fscanf(reader, "%d", &data)

		if (scanErr != nil) && (scanErr != io.EOF) {
			fmt.Printf("\n\nScan Error: %v\n\nExiting...\n\n", scanErr)
			return
		}

		if scanErr != io.EOF {
			fmt.Printf("\n%d", data)

			totalSuccessfulParses += successfulParses
		}
	}

	fmt.Printf("\n\nTotal number of successful parses: %d. The fscanf function of the fmt library returns the number of successful parses, as well as an error status.\n\n", totalSuccessfulParses)

	closeErr = reader.Close()
	if closeErr != nil {
		fmt.Printf("\n\nClose Error: %v\n\nExiting...\n\n")
		return
	}
}

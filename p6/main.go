/*
Problem #5: The fscanf Function
Author: Corey Prak
Date Created: 04/03/2014
Comments:
Had to refer to documentation so much. Really not used to the organization:
func Open(name string) (file *File, err error)

This is how I see it, hope it helps:
func *function name*(*argument name and data type*) (*return parameters, name then datatype*)
*/

package main

//importing fmt for file I/O and os for opening file. 
import "fmt"
import "os"

func main(){        
  // Open, used for reading,
  // Takes a string, returns a pointer to a file (file) & error (err)
  filePtr, fErr := os.Open("testdata6")
  	// If fErr is not nil, then it'll be a *PathError.
    if fErr != nil {
			fmt.Println(fErr)
    }

  // Dynamically create a slice of 100 bytes.
	data := make([]byte, 100)

  // Read the file. Number of bytes depends on the size of the slice. 
  // count will receive Read's return value of the number of bytes read. 
	count, cErr := filePtr.Read(data)
  // If xErr is not nil, then it'll be a *PathError. 
  if cErr != nil {
		fmt.Println(fErr)
  }

  // Print the number of bytes read, as well as the data.
  // I can't find any documentation on wth data[] is. 
  fmt.Printf("read %d bytes: %q\n", count, data[:count])
}

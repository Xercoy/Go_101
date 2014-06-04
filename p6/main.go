/*******************************************************************************
Problem #5: The fscanf Function
Author: Corey Prak
Date Created: 04/03/2014
Comments:
06/04/2014

Added comments

Changed code around
- Removed short declarations.
- Closed the file. 
*******************************************************************************/

package main

// Importing fmt for file I/O and os for opening file. 
import "fmt"
import "os"

func main(){ 

  /* Declare filtPtr to be a  pointer to a File. File is a type defined in the 
     os package. */
  var filePtr *os.File
  var fErr, cErr error

  var count int

  /* Declare data to be a slice of bytes, then allocate contiguous space of 100
     bytes for the data slice to refer to. */
  var data []byte
  data = make([]byte, 100)

  /* Call the function Open from the os package, provide the file to be opened 
     as an argument. If successful, return an address which filePtr will be 
     assigned to and leave fErr untouched with its default value of nil. If
     the file is unable to be opened, then an fErr will not be nil, but another
     error instead. */
  filePtr, fErr = os.Open("testdata6")
 	// If fErr is not nil, then it'll be a *PathError.
    if fErr != nil {
			fmt.Println(fErr)
    }

  // Read the file. Number of bytes depends on the size of the slice. 
  // count will receive Read's return value of the number of bytes read. 
	count, cErr = filePtr.Read(data)
  // If xErr is not nil, then it'll be a *PathError. 
  if cErr != nil {
		fmt.Println(fErr)
  }

  // Print the number of bytes read, as well as the data.
  // I can't find any documentation on wth data[] is. 
  fmt.Printf("read %d bytes: %q\n", count, data[:count])

  /* Close the file. ALWAYS REMEMBER THIS! This is not a general function of 
     the os package, but instead a method which belongs to a pointer of to a File
     (*File), which must be accessed as such, similar to what we've seen when 
     accessing the functions of a package. */
  filePtr.Close()
}

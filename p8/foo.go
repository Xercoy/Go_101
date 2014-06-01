package main

import(
	"fmt"
  "os"
  "strconv"
)

func main() {
  data := make([]byte, 10)
  
  // Iterator for the for loop.
  i := 0;
  count := 0;

  // Open the file, return address to filePtr, fErr if possible.
	filePtr, fErr := os.Open("testdata8")  
  if fErr != nil {
		fmt.Println(fErr)
  }

  count, scanErr := filePtr.Read(data)
  if scanErr != nil {
    fmt.Println(scanErr)
  }
  
  fmt.Printf("Read in the number %q. in %d bytes.\n", data[:count], count)
  fileNumberValue, fNumErr := strconv.Atoi(string(data[:count]))
  if fNumErr != nil {
    fmt.Println(fNumErr)
  }
  

  testValue := data[7]
  fmt.Println(testValue)
  
  fmt.Printf("the type of fileNumberValue is %T while i is of type %T.\n", fileNumberValue, i)

  fmt.Printf("i is %d while fileNumberValue is %d.\n", i, fileNumberValue)
  if i < fileNumberValue {
    fmt.Println("i < fileNumberValue = true")
  } else {
    fmt.Println("i < fileNumberValue = false")
  }
 

  for i = 0; i < fileNumberValue  ; i = i+1 {
	  fmt.Printf("*")
  }
  
		fmt.Printf("\n")

  fmt.Printf("data[:count] is of type %T, data[count] is of type %T.\n", data[:count], data[count])

  fmt.Println(data) 
}

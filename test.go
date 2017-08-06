package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
  var fileName string
  fmt.Printf("Enter file name:")
  fmt.Scanf("%s", &fileName)
  file, err := ioutil.ReadFile(fileName)

   if err != nil {
       fmt.Println("file not found")
       return
   }

   str := string(file)
   fmt.Println(strings.Count(str, "\n"))
}

package main

import (
  "eshop/translate"
  "fmt"
)

func main() {
  s, err:=translate.Translate("hello", "mn")
  if err!=nil {
    fmt.Printf("%v", err)
  } else {
    fmt.Printf("%s", s)  
  }

}

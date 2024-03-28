package main

import (
	"main/list"
  "fmt"
)

func main() {
	llist := list.ArrayList{}

  llist.Initialize(5)
  llist.Add(5)
  llist.Add(10)
  llist.Add(15)

  fmt.Println(llist.Size())

  

  for i:=0; i<llist.Size(); i++{

    val, err := llist.Get(i)

    if err != nil{
      fmt.Println(err)
    }else{
      fmt.Print("entrei aqui")
      fmt.Println(val)
    }
  }
}

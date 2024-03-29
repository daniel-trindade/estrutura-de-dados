package main

import (
	"main/list"
  "fmt"
)

func main() {
	llist := list.LinkedList{}

  llist.Add(5)
  llist.Add(10)
  llist.Add(15)
  llist.Add(20)
  llist.Add(25)
  llist.Add(13)
  llist.Add(22)

  fmt.Println("O número de valores no vetor é:", llist.Size(), "são eles:")

  for i:=0; i<llist.Size(); i++{

    val, err := llist.Get(i)

    if err != nil{
      fmt.Println(err)
    }else{
      fmt.Print(val, ", ")
    }
    
  }
  fmt.Println()
  fmt.Println("*@ testando adicionar no index @*")

  err := llist.AddOnIndex(55, 5)
  if err != nil{
    fmt.Println(err)
  }

  fmt.Println("O número de valores no vetor é:", llist.Size(), "são eles:")

  for i:=0; i<llist.Size(); i++{

    val, err := llist.Get(i)

    if err != nil{
      fmt.Println(err)
    }else{
      fmt.Print(val, ", ")
    }
    
  }

  fmt.Println()
  fmt.Println("*@ testando remover no index @*")

  erro := llist.RemoveOnIndex(5)
  if erro != nil{
    fmt.Println(erro)
  }

  fmt.Println("O número de valores no vetor é:", llist.Size(), "são eles:")

  for i:=0; i<llist.Size(); i++{

    val, err := llist.Get(i)

    if err != nil{
      fmt.Println(err)
    }else{
      fmt.Print(val, ", ")
    }

  }

  fmt.Println()
  fmt.Println("*@ testando set no index @*")

  errora := llist.Set(18, 7)
  if errora != nil{
    fmt.Println(errora)
  }

  fmt.Println("O número de valores no vetor é:", llist.Size(), "são eles:")

  for i:=0; i<llist.Size(); i++{

    val, err := llist.Get(i)

    if err != nil{
      fmt.Println(err)
    }else{
      fmt.Print(val, ", ")
    }

  }
}

  

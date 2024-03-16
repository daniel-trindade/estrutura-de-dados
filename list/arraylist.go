package list

import (
  "errors"
  "fmt"
)

type ArrayList struct{
  values []int
  inserted int
}

//Função para Inicializar valores

func (arraylist *ArrayList) Initialize(size int) error {
  if size > 0{
    arraylist.values = make([]int, size)
    fmt.Println("Array criado com tamanho ", arraylist.inserted)
  } else {
    return errors.New("O tamanho ptretendido precisa ser maior que 0 (Zero)")
  }
  return nil
}

//Função para dobrar tamanho do array

func (arraylist *ArrayList) DoubleSize(){
  arraydoubled := make([]int, 2*len(arraylist.values))
  for i := 0; i<len(arraylist.values); i++{
    arraydoubled[i] = arraylist.values[i]
  }
  arraylist.values = arraydoubled
  fmt.Println("O Array teve seu tamanho dobrado. Tamanho atual: ", arraylist.inserted)
}

func (arraylist *ArrayList) Add (value int){
  if(len(arraylist.values) > arraylist.inserted){
    arraylist.values[arraylist.inserted] = value
  }else{
    arraylist.DoubleSize()
    arraylist.values[arraylist.inserted] = value
  }
}

func (arraylist *ArrayList) AddOnIndex (value, index int) error {
  
}
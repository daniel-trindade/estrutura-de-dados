package list

import (
	"container/list"
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
    fmt.Println("Valor ", value, " inserido com Sucesso!")
  }else{
    arraylist.DoubleSize()
    arraylist.values[arraylist.inserted] = value
    fmt.Println("Array redimencionado e valor ", value, " inserido com Sucesso!")
  }
}

func (arraylist *ArrayList) AddOnIndex (value, index int) error {
  if arraylist.inserted == len(arraylist.values){
    arraylist.DoubleSize()
  }
  if index > arraylist.inserted{
    return errors.New("O index para inserção do valor precisa ser menor que a quantidade de números já inseridos!")
  }
  arrayAux := make([]int, len[arraylist.values])
  for i := 0; i<index; i++{
    arrayAux[i]=arraylist.values[i]
  }
}
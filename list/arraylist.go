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

//Função para adicionar valor no array
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
 //Função para adicionar valor no array em um índice específico
func (arraylist *ArrayList) AddOnIndex (value, index int) error {
  if arraylist.inserted == len(arraylist.values){
    arraylist.DoubleSize()
  }
  if index > arraylist.inserted || index <0{
    return errors.New("O index para inserção do valor precisa estar entre 0 e a quantidade de números já inseridos!")
  }
  arrayAux := make([]int, len(arraylist.values))
  for i := 0; i<index; i++{
    arrayAux[i]=arraylist.values[i]
  }
  
  arrayAux[index] = value
  
  for i:=index; i<arraylist.inserted+1; i++{
    arrayAux[i+1]=arraylist.values[i]
  }

  arraylist.values = arrayAux

  arraylist.inserted++

  return nil
}

//Função para remover um valor em um índice específico
func (arraylist *ArrayList) RemoveOnIndex(index int) error {
  if index > arraylist.inserted || index <0{
    return errors.New("O index para remoção do valor precisa estar entre 0 e a quantidade de números já inseridos!")
  }else{
    for i:= index; i<arraylist.inserted-1; i++{
      arraylist.values[i] = arraylist.values[i+1]
    }
  
    arraylist.inserted--
  
    return nil
  }
}

//Função para retornar um valor de um índice específico
func (arraylist *ArrayList) Get(index int) (int, error){
  if index > arraylist.inserted || index < 0 {
    return -1, errors.New("O index do valor precisa estar entre 0 e a quantidade de números já inseridos!")
  }else{
    return arraylist.values[index], nil
  }
}

//Função para editar um valor em um indice específico
func (arraylist *ArrayList) Set(value int, index int) error{
  if index > arraylist.inserted || index < 0{
    return errors.New("O index precisa estar entre 0 e a quantidade de números já inseridos!")
  }else{
    arraylist.values[index] = value
  }
}


//Função para retornar o tamanho do array
func (arraylist *ArrayList) Size() int{
  return arraylist.inserted
}
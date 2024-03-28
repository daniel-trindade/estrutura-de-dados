package list

import (
  "errors"
  "fmt"
)

type Node struct{
  value int
  next *Node
}

type LinkedList struct{
  head *Node
  inserted int
}

//Função para Inserir valores
func (linkedlist *LinkedList) Add(value int){

  newNode := Node{value, nil}

  aux := linkedlist.head
  prev := aux

  if linkedlist.head == nil{
    
    linkedlist.head = &newNode
    
  }else{

    for aux != nil {
      prev = aux
      aux = aux.next
    }

    prev.next = &newNode
  }

   linkedlist.inserted++

}

//Função para Inserir valores em um índice específico

func (linkedlist *LinkedList) AddOnIndex(value, index int) error{

 if index > linkedlist.inserted || index < 0 {
    return errors.New("O index para inserção do valor precisa estar entre 0 e a quantidade de números já inseridos")
 }else{
    newNode := Node{value, nil}

    if index == 0{
      newNode.next = linkedlist.head
      linkedlist.head = &newNode

      return nil
    }
   
    aux := linkedlist.head
    prev := aux

    for i:=0; i<index; i++{
      prev = aux
      aux = aux.next
    }

    prev.next = &newNode
    newNode.next = aux
   
    linkedlist.inserted++
  }

}

//Função para remover um valor de um index em específico
func (linkedlist *LinkedList) RemoveOnIndex(index int) error{
  if index > linkedlist.inserted || index < 0{
    errors.New("O index do valor precisa estar entre 0 e a quantidade de números já inseridos")
  }else{
    aux := linkedlist.head
    prev := aux
    
    if index == 0{
      linkedlist.head = aux.next
      return nil
    }
    
    for i:=0; i<index; i++{
      prev = aux
      aux = aux.next
    }

    prev.next = aux.next

    linkedlist.inserted--
  }

  return nil
}

//Função para retornar um valor de um indice específico
func (linkedlist *LinkedList) Get(index int) (int, error){

  if index > linkedlist.inserted || index < 0{
    return -1, errors.New("O index do valor precisa estar entre 0 e a quantidade de números já inseridos")
  }else{
    if index == 0{
      return linkedlist.head.value, nil
    }

    aux := linkedlist.head
    
    for i:=0; i<index; i++{
      aux = aux.next
    }

    return aux.value, nil
  }
  
}
package list

import (
  "errors"
)

type Nod struct{
  prev *Nod
  value int
  next *Nod
}

type DoublyLinkedList struct{
  head *Nod
  tail *Nod
  inserted int
}

//Função para Inserir valores
func (doublylinkedlist *DoublyLinkedList) Add(value int){

  newNod := Nod{nil, value, nil}

  if doublylinkedlist.inserted > 0{
  
    doublylinkedlist.tail.next = &newNod
    newNod.prev = doublylinkedlist.tail
    doublylinkedlist.tail = &newNod

  }else{
    
    doublylinkedlist.head = &newNod
    doublylinkedlist.tail = &newNod
    
  }
  
    doublylinkedlist.inserted++

}

//Função para Inserir valores em um índice específico
func (doublylinkedlist *DoublyLinkedList) AddOnIndex(value, index int) error{
  
  if index >= doublylinkedlist.inserted || index < 0{
    
    return errors.New("O index para inserção do valor precisa estar entre 0 e a quantidade de números já inseridos")
    
  }else{
    
    newNod := Nod{nil, value, nil}
    
    if index == 0{
      
      newNod.next = doublylinkedlist.head
      doublylinkedlist.head.prev = &newNod
      doublylinkedlist.head = &newNod
      
    }else if index == doublylinkedlist.inserted{
      
      newNod.prev = doublylinkedlist.tail
      doublylinkedlist.tail.next = &newNod
      doublylinkedlist.tail = &newNod
      
    }else{
      
      aux := doublylinkedlist.head

      for i:=0; i<index; i++{
        aux = aux.next
      }

      newNod.next = aux
      newNod.prev = aux.prev
      aux.prev.next = &newNod
      aux.prev = &newNod
      
    }
  }

  doublylinkedlist.inserted++

  return nil
}

//Função para remover um valor de um índice específico
func (doublylinkedlist *DoublyLinkedList) RemoveOnIndex(index int) error{
  if index >= doublylinkedlist.inserted || index < 0{
    
    return errors.New("O index do valor precisa estar entre 0 e a quantidade de números já inseridos")
    
  }else{

    if index == 0{
      
      doublylinkedlist.head = doublylinkedlist.head.next
      doublylinkedlist.head.prev = nil
      
    }else if index == doublylinkedlist.inserted-1{

      doublylinkedlist.tail = doublylinkedlist.tail.prev
      doublylinkedlist.tail.next = nil
      
    }else{

      aux := doublylinkedlist.head

      for i:=0; i<index; i++{
        aux = aux.next
      }

      aux.prev.next = aux.next
      aux.next.prev = aux.prev
      
    }
    
  }

  doublylinkedlist.inserted--

  return nil
}

//Função para retornar um valor de um índice específico
func (doublylinkedlist *DoublyLinkedList) Get(index int) (int, error){

  if index > doublylinkedlist.inserted || index < 0{
    
    return -1, errors.New("O index do valor precisa estar entre 0 e a quantidade de números já inseridos")
    
  }else{
    
    if index == 0{
      
      return doublylinkedlist.head.value, nil
      
    }else if index == doublylinkedlist.inserted-1{

      return doublylinkedlist.tail.value, nil
      
    } else{

      aux := doublylinkedlist.head

      for i:=0; i<index; i++{
        aux = aux.next
      }

      return aux.value, nil
      
    }
  }
}

//Função para editar um valor em um índice específico{
func (doublylinkedlist *DoublyLinkedList) Set(value, index int) error{
  if index > doublylinkedlist.inserted || index < 0{

    return errors.New("O index do valor precisa estar entre 0 e a quantidade de números já inseridos")

  }else{

    if index == 0{

      doublylinkedlist.head.value = value

      return nil

    }else if index == doublylinkedlist.inserted-1{

      doublylinkedlist.tail.value = value

      return nil

    }else{

      aux := doublylinkedlist.head

      for i:=0; i<index; i++{
        aux = aux.next
      }

      aux.value = value

      return nil

    }
  }
}

//Função para retornar o tamanho da linkeList
func (doublylinkedlist *DoublyLinkedList) Size() int{
  return doublylinkedlist.inserted
}

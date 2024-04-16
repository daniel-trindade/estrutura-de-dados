package sorting


func SelectionSort(values []int){
  for i:=0; i<len(values)-1; i++{
    iMin := i
    for j:=i+1; j<len(values); j++{
      if values[j] < values[iMin]{
        iMin = j
      }
    }
    values[i], values[iMin] = values[iMin], values[i]
  }
}
package main

import (
  "fmt"
  "time"
)

func getData() []int{
  return []int{4,3,2,6,1,13,2,3,2,6,7,8,0,9,12,10,11,15,14,0}
}

func main() {
  start := time.Now()
  array := getData();
  array = mergeSort(array);
  fmt.Println(array);
  fmt.Println(time.Now().Sub(start));
}

func mergeSort(array []int) []int{
  length := len(array);

  if length <= 1 {
    return array
  }

  middle := int(length / 2);

  left := array[0: middle]
  right := array[middle: length]

  sortedLeft := mergeSort(left)
  sortedRight := mergeSort(right)

  return merge(sortedLeft, sortedRight)
}

func merge(left []int, right []int) []int{
  result := make([]int, len(left) + len(right))

  i := 0

  for len(left) > 0 && len(right) > 0 {
    if left[0] < right[0] {
      result[i] = left[0]
      left = left[1:]
    } else {
      result[i] = right[0]
      right = right[1:]
    }
    i++;
  }

  for j := 0; j < len(left); j++ {
    result[i] = left[j];
    i++;
  }

  for j := 0; j < len(right); j++ {
    result[i] = right[j];
    i++;
  }
  
  return result;
}


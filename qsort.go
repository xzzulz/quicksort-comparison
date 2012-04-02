// quicksort in go
// 
// An non optimal imploementation of quicksort
// in go lang, made for comparative purposes of this
// project.
//

package main

import (
	"fmt"
)

func qsort( list []int ) []int {

  if len(list) == 0 { return list }
  
  num := list[0]
  list = list[1:]

  lesser_filter := func (n int) bool { if n<=num { return true }; return false }  
  greater_filter := func (n int) bool { if n>num { return true }; return false }
  
  lesser := filter( list, lesser_filter )
  greater := filter( list, greater_filter )
  
  return concat( append( qsort( lesser ), num ) , qsort( greater ) )
}



type filter_func func(n int) bool


func filter( list []int, f filter_func ) []int {
  res := make( []int, 0, len(list) )
  for _, pik := range list {
    if f(pik) { res = append(res,pik) }
  }
  return res
}

func concat( arr1, arr2 []int ) []int { 
  conc := make( []int, len( arr1 ) + len( arr2 ) ) 
  copy( conc, arr1 ) 
  copy( conc[ len(arr1): ], arr2 ) 
  return conc 
}


var arr = []int{ 9, 3, 5, 15, 1, 7, 11, 13, 6 }
  
func main() {
	
  arr_sorted :=  qsort( arr )
  
  fmt.Println( arr )
  fmt.Println( arr_sorted )
  
}




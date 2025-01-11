package main

import "fmt"

func checkEven(n int) bool {
  if n%2==0{
    return true
  } else{
  return false
  }
}

func main(){
  fmt.Println("Input a number:")
  var num int
  fmt.Scan(&num)
  for i:=0;i<=num;i++{
    if checkEven(i){
      fmt.Println(i," ")
    }
  }
  return 
}

package main

import "fmt"

func main(){
  
  var a int = 5
  var b int = 8

  var opt int 

  for i:=0;i<=4;i++ {
    fmt.Scan(&opt);
    switch opt {
    case 1:
      fmt.Println(a+b)
    case 2:
      fmt.Println(a-b)
    case 3:
      fmt.Println(a*b)
    case 4:
      fmt.Println(b/a)
    }
  }
  return 
}

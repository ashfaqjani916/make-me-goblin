package main

import (
	"fmt"
	"time"
)


func main()  {

	data := make(chan string)
	
go func()  {
	time.Sleep(time.Second * 2)
		data <- "this is goroutine One"
}()


go func ()  {
	time.Sleep(time.Second * 2)
		data <- "this is goroutine two"
}()

	select{
	case msg1:= <-data:
	fmt.Println(msg1)
	case msg2 := <-data:
	fmt.Println(msg2)
	case <- time.After(5*time.Second):
	fmt.Println("Request timed out")
	}

}

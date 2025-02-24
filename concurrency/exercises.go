package main

import (
	"fmt"
	"sync"
	"time"
)



func ExerciseOne()  {

	channel := make(chan string,3)
	var waitGroup sync.WaitGroup


	go func() {
	defer waitGroup.Done()
		fmt.Println("this is from R1")
		fmt.Println(<-channel)
	}()


	go func() {
	defer waitGroup.Done()
	fmt.Println("this is from R2")
		fmt.Println(<-channel)
	}()

	go func() {
		defer waitGroup.Done()
		time.Sleep(time.Millisecond * 100)
	fmt.Println("this is from R3")
		fmt.Println(<-channel)	
	}()

		channel <- "this is the data that is being put"	
		channel <- "this is the second sentence"
		channel <- "this is the 3rd one"
		close(channel)

	waitGroup.Add(3)
	waitGroup.Wait()

}

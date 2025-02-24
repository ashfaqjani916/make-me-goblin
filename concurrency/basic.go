package main

import (
	"fmt"
)

 func smtg()  {
	fmt.Println("ash")
}

func Anothermain()  {

	// before calling a function when we specify "go" then is it considered as a goroutine and a seperate goroutine will branch out from the main execution. this is basically the fork join model. the brached out threaad after executing will join the main branch.  

	myChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}()

	anotherChannel := make(chan string)
	go func() {
		anotherChannel <- "cats and dogs"	
	}()

// In order for the goroutines to communicate with eachother, we will be using channels

// **channels** are basically FIFO queues 

// so the channel waits the execution of the main goroutine till the data we are requested has arrived from the child 
// goroutine.

/*
	Now if we have more than one goroutine and we have to select the data from there, we use the select statement

*/


	// now we are selecting the data from either of the channel 
	select {
	case msgfromChannel := <- myChannel:
		fmt.Println(msgfromChannel)
	case msgfromAnotherChannel := <-anotherChannel:
		fmt.Println(msgfromAnotherChannel)
	}

	pipeline()
	// here the msg that arrives first gets printed. 

}



func pipeline() {
	//here we will have a slice of numbers and we will find the square of each number using the pipeline
	arr := []int{2,3,5,7}

	arr = append(arr, 19)

	for _,n := range arr {
		fmt.Println(n)
	}

	//stage 1 to put all of them into a channel

	//stage 2 to find the squares 

	//
}

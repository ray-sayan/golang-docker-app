package main

import (
	"fmt"
)

func receiver1(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
}

func receiver2(ch chan int) {
	for i := 6; i <= 10; i++ {
		ch <- i
	}
}

func receiver3(ch chan int) {
	for i := 11; i <= 15; i++ {
		ch <- i
	}
}

func sender(ch chan int) {
	for i := 0; i < 15; i++ {
		fmt.Println(<-ch)
	}
}

// func sender(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for i := 0; i < 15; i++ {
// 		fmt.Println(<-ch)
// 	}
// }

// using 3 channels
// read from 3 channels using select
// func sender(ch1, ch2, ch3 chan int) {

// 	for i := 0; i < 15; i++ {

// 		select {

// 		case v := <-ch1:
// 			if v != 0 {
// 				fmt.Println(v)
// 			}

// 		case v := <-ch2:
// 			if v != 0 {
// 				fmt.Println(v)
// 			}

// 		case v := <-ch3:
// 			if v != 0 {
// 				fmt.Println(v)
// 			}
// 		}
// 	}
// }

func main() {

	ch := make(chan int)

	go func() {
		receiver1(ch)
		receiver2(ch)
		receiver3(ch)
	}()

	sender(ch)

	//using 3 channels
	// ch1 := make(chan int)
	// ch2 := make(chan int)
	// ch3 := make(chan int)

	// go receiver1(ch1)
	// go receiver2(ch2)
	// go receiver3(ch3)

	// sender(ch1, ch2, ch3)

	//oddeven task
	// oddChan := make(chan int)
	// evenChan := make(chan int)

	// go oddeven.Odd(oddChan)
	// go oddeven.Even(evenChan)

	// oddeven.Receiver(oddChan, evenChan)
}

package oddeven

import "fmt"

// Odd goroutine sends odd numbers
func Odd(oddChan chan int) {
	oddChan <- 1
	oddChan <- 3
	oddChan <- 5
	oddChan <- 7
}

// Even goroutine sends even numbers
func Even(evenChan chan int) {
	evenChan <- 2
	evenChan <- 4
	evenChan <- 6
	evenChan <- 8
}

// Receiver prints values in sequence
func Receiver(oddChan chan int, evenChan chan int) {

	// receive odd
	fmt.Println(<-oddChan)
	// receive even
	fmt.Println(<-evenChan)

	fmt.Println(<-oddChan)
	fmt.Println(<-evenChan)

	fmt.Println(<-oddChan)
	fmt.Println(<-evenChan)

	fmt.Println(<-oddChan)
	fmt.Println(<-evenChan)
}

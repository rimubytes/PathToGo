func printMessage() {
	fmt.Println("Hello, Goroutine")
}

func main() {
	go printMessage() // You can create a Goroutine by adding the go keyword before a function call
	time.Sleep(1 * time.Second) // Ensuring the Goroutine finishes before the program exits
}
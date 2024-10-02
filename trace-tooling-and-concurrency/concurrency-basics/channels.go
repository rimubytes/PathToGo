//Channels are used to communicate between goroutines safely

func sendMessage(ch chan string) {
	ch <- "Hello from Goroutine"
}

func main() {
	ch := make(chan string)
	go sendMessage(ch)
	msg := <-ch
	fmt.Println(msg)
}

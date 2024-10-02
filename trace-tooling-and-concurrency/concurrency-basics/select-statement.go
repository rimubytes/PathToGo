// select statement allows you to wait on multiple channel operations

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {time.Sleep(1 * time.Second); ch1 <- "From ch1"}()
	go func() {time.Sleep(1 * time.Second); ch2 <- "From ch2"}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
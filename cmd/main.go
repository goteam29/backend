package main

import "fmt"

func main() {
	fmt.Println("server started!")
	ch := make(chan int)
	<-ch
}

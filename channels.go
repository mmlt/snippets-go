package main

// Show how a single channel can be used to stop a go routine and wait for its exit.

import "fmt"
import "time"


func main() {
	fmt.Println("start")
	quit := make(chan string)
	
	go func() {
		for {
			<- time.After(time.Second * 1)
			select {
			case <- quit:
				fmt.Println("start quit")
				<- time.After(time.Second * 1)
				fmt.Println("exit quit")
				quit <- "done"
				return
			}
		}
	}()
	
	quit <- "1"
	<- quit
	fmt.Println("exit")	
}


/*
package main

import "fmt"

func main() {
	fmt.Println("start")
	quit := make(chan string, 1)
	
	quit <- "1"
	<- quit
	fmt.Println("exit")	
}*/
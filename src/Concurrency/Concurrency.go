package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//sample1()
	//sample2()
	//sample3()
	//sample4()
	//sample5()
	//sample6()
	//sample7()
	sample8()
}

func say(message string) {
	for i := 0; i < 5; i++ {
		fmt.Println(message)
		time.Sleep(time.Millisecond*100)
	}
}



func sample1() {
	say("Hello")
	say("World")
}

func sample2() {
	go say("Hello")
	say("World")
}

func sample3() {
	say("Hello")
	go say("World")
}

func sample4() {
	go say("Hello")
	go say("World")
}

func sample5() {
	go say("Hello")
	go say("World")

	time.Sleep(time.Second * 2)
}

var waitGroup sync.WaitGroup

func sample6() {

	waitGroup.Add(1)
	go sayAndNotifyWaitGroup("Hello")
	waitGroup.Add(1)
	go sayAndNotifyWaitGroup("World")
	waitGroup.Wait()
}

func sayAndNotifyWaitGroup(message string) {
	for i := 0; i < 5; i++ {
		fmt.Println(message)
		time.Sleep(time.Millisecond*100)
	}

	waitGroup.Done()
}


func sample7(){
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go sayAndWait(done)

	<-done // Ignore return.
	fmt.Println("Finally!")
}

func sayAndWait(done chan bool) {
	fmt.Println("Hello")
	time.Sleep(4 * time.Second)
	fmt.Println("World")
	done <- true
}

func sample8(){
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go sayAndWait(done)

	isWaiting := true

	for isWaiting {
		select {
		case <-done:
			fmt.Println("FINALLY DONE")
			isWaiting = false
		default:
			fmt.Println("Wait for it!")
			time.Sleep(time.Millisecond*500)
		}
	}

	fmt.Println("Finally!")
}
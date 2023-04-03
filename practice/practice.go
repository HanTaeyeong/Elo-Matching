package practice

import (
	"fmt"
	"time"
)

func putString(channel chan int, num int) {
	channel <- num
}
func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("do")
		}
	}
}

func DoneChannelPattern() {
	doneChannel := make(chan bool)
	go doWork(doneChannel)
	time.Sleep(time.Second * 2)
	close(doneChannel)
	fmt.Println("finished")
}

func Practice() {
	myChannel := make(chan int, 100)

	for i := 0; i < 10; i += 1 {
		go putString(myChannel, i)
	}
	time.Sleep(time.Second * 2)

	for msg := range myChannel {
		fmt.Println(msg)
	}

}

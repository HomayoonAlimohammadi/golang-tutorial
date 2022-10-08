package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
)

var DoneChannel chan bool

type Task struct {
	Id string
}

func consume(taskChan chan Task) {
	for task := range taskChan {
		// do job
		fmt.Printf("Job %s started\n", task.Id)
		time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
		fmt.Printf("Job %s finished\n", task.Id)
	}
}

func dispatch(taskChan chan Task, stop chan bool) {
	go consume(taskChan)
	go consume(taskChan)
	go consume(taskChan)

	<-stop
}

func produce() chan Task {
	taskChan := make(chan Task, 10)
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			taskChan <- Task{
				Id: uuid.New().String(),
			}
		}
	}()
	return taskChan
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println("Signal received. Program is going to stop...")
		fmt.Println(sig)
		DoneChannel <- true
	}()
	taskChan := produce()
	dispatch(taskChan, DoneChannel)

	fmt.Println("Program is going to close")

}

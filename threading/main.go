package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 1; i < 10; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(time.Second)
	}
}

func main_old_1() {
	fmt.Println("Olá Mundo")

	go task("tarefa 1") // T2 // go routine
	go task("tarefa 2") // T3 // go routine
	task("tarefa 3") // T1
}

func main_old_2() {
	canal := make(chan string) // canal de comunicação entre as threads
	
	// thread 2
	go func() {
		canal <- "Olá mundo!"

		canal <- "Olá mundo 2!"
	}()

	// thread 1
	msg := <- canal
	fmt.Println(msg)

	fmt.Println(<- canal)
}

func publish(ch chan int) {
	for i := 1; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func reader(canal chan int) {
	for x := range canal {
		fmt.Println(x)
	}
}

func main_old_3() {
	canal := make(chan int)
	go publish(canal)
	go reader(canal)
	time.Sleep(time.Second * 5)
}

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d \n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int)

	qtsWorkers := 100

	// inicializa os workers
	for i := 0; i < qtsWorkers; i++ {
		go worker(i, ch)
	}

	// joga carga para os works
	for i := 0; i < 100000; i++ {
		ch <- i
	}
}

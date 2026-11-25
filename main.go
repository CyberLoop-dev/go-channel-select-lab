package main

import (
	"flag"
	"fmt"
	"time"
)

func buffered() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20
	fmt.Printf("buffered: len=%d cap=%d\n", len(ch), cap(ch))
	fmt.Printf("received %d then %d\n", <-ch, <-ch)
}

func unbuffered() {
	ch := make(chan string)
	done := make(chan struct{})
	go func() {
		value := <-ch
		fmt.Println("receiver got:", value)
		close(done)
	}()
	ch <- "hello from sender"
	<-done
	fmt.Println("unbuffered send completed after a receiver was ready")
}

func selectDemo() {
	ready := make(chan string, 1)
	ready <- "message ready"
	select {
	case value := <-ready:
		fmt.Println("select received:", value)
	case <-time.After(50 * time.Millisecond):
		fmt.Println("select timed out")
	}
	empty := make(chan int)
	select {
	case value := <-empty:
		fmt.Println("unexpected value:", value)
	default:
		fmt.Println("default branch: no receive was ready")
	}
}

func closed() {
	ch := make(chan int, 1)
	ch <- 7
	close(ch)
	value, ok := <-ch
	fmt.Printf("first receive: value=%d ok=%t\n", value, ok)
	value, ok = <-ch
	fmt.Printf("after drain: value=%d ok=%t\n", value, ok)
}

func main() {
	scenario := flag.String("scenario", "all", "buffered, unbuffered, select, closed, or all")
	flag.Parse()
	demos := map[string]func(){"buffered": buffered, "unbuffered": unbuffered, "select": selectDemo, "closed": closed}
	if *scenario == "all" {
		for _, name := range []string{"buffered", "unbuffered", "select", "closed"} {
			fmt.Println("===", name, "==="); demos[name]()
		}
		return
	}
	demo, ok := demos[*scenario]
	if !ok { fmt.Println("unknown scenario:", *scenario); return }
	demo()
}

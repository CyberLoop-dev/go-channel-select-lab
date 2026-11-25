package main

import "testing"

func TestBufferedChannelCapacity(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	if len(ch) != 2 || cap(ch) != 2 { t.Fatalf("got len=%d cap=%d", len(ch), cap(ch)) }
}

func TestClosedChannelReportsFalse(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 9
	close(ch)
	<-ch
	_, ok := <-ch
	if ok { t.Fatal("receive from drained closed channel should report ok=false") }
}

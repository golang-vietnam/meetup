package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	grpc "google.golang.org/grpc"
)

var numWorker = 50
var numRequest = 10000
var successCounter uint64

var hClient HelloClient

var p = &Payload{
	A: "12345678901234567890",
	B: "abcdefghijklmnopqus",
	C: 43545,
	D: "asgkasjeglkajsgjalsjdglaskdjglads",
	E: true,
	F: -135465748,
	G: "gejoiwbbbdfasaaaaaaaaaaaaaaaaa",
}

func main() {
	ch := make(chan struct{}, numWorker)

	// worker
	var wg sync.WaitGroup
	for i := 0; i < numWorker; i++ {
		wg.Add(1)
		go func() {
			for range ch {
				doTask()
			}
			wg.Done()
		}()
	}

	// init client
	conn, err := grpc.Dial("127.0.0.1:9999", grpc.WithInsecure())
	if err != nil {
		return
	}
	hClient = NewHelloClient(conn)

	time.Sleep(2 * time.Second)

	start := time.Now()
	for i := 0; i < numRequest; i++ {
		ch <- struct{}{}
	}
	close(ch)
	wg.Wait()
	end := time.Now()

	fmt.Println(end.Sub(start).Seconds(), fmt.Sprintf("%d / %d", successCounter, numRequest))
}

func doTask() {
	_, err := hClient.SayHi(context.Background(), p)
	if err != nil {
		return
	}

	atomic.AddUint64(&successCounter, 1)
}

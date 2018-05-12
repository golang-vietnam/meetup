package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/golang/protobuf/proto"
)

var numWorker = 50
var numRequest = 1000000
var successCounter uint64

var tr = &http.Transport{
	MaxIdleConnsPerHost: 200,
}
var httpClient = http.Client{Transport: tr}

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
	pl, err := proto.Marshal(p)
	if err != nil {
		return
	}

	resp, err := httpClient.Post("http://127.0.0.1:8888/hello", "application/protobuf", bytes.NewReader(pl))
	if err != nil {
		return
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var tmp Payload
	if err = proto.Unmarshal(buf, &tmp); err != nil {
		return
	}

	atomic.AddUint64(&successCounter, 1)
}

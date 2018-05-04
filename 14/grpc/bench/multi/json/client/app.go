package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

var numWorker = 50
var numRequest = 1000000
var successCounter uint64

var tr = &http.Transport{
	MaxIdleConnsPerHost: numWorker,
}
var httpClient = http.Client{Transport: tr}

type payload struct {
	A string `json:"a"`
	B string `json:"b"`
	C int    `json:"c"`
	D string `json:"d"`
	E bool   `json:"e"`
	F int64  `json:"f"`
	G string `json:"g"`
}

var p = &payload{
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
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(p); err != nil {
		fmt.Println(err)
		return
	}

	resp, err := httpClient.Post("http://127.0.0.1:8888/hello", "application/json; charset=utf-8", b)
	if err != nil {
		fmt.Println(err)
		return
	}

	var tmp payload
	if err = json.NewDecoder(resp.Body).Decode(&tmp); err != nil {
		return
	}

	atomic.AddUint64(&successCounter, 1)
}

package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	etcdClient "github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	grpc "google.golang.org/grpc"
)

const (
	serviceName = "helloService"
)

var numWorker = 100
var numRequest = 1000000
var hClient HelloClient
var successCounter uint64

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

	// get etcd client
	cli, _ := getETCDClient()

	// init balancer
	balancer := grpc.RoundRobin(&etcdnaming.GRPCResolver{Client: cli})
	conn, err := grpc.Dial(serviceName, grpc.WithInsecure(), grpc.WithBalancer(balancer))
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

func getETCDClient() (*etcdClient.Client, error) {
	return etcdClient.New(etcdClient.Config{
		Endpoints: strings.Split("http://127.0.0.1:2379", ","), // etcd endpoints
	})
}

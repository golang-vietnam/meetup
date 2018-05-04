package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	etcdClient "github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	grpc "google.golang.org/grpc"
	naming "google.golang.org/grpc/naming"
)

const (
	publicIP    = "127.0.0.1"
	publicPort  = 9903
	serviceName = "helloService"
)

func main() {
	// os signal handling
	sigs := make(chan os.Signal, 10)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGTSTP, syscall.SIGEMT)

	// get etcd client
	cli, _ := getETCDClient()

	// try to listening to bind address
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", publicPort))
	if err != nil {
		return
	}
	s := grpc.NewServer()
	RegisterHelloServer(s, &helloServer{})

	go func() {
		if e := s.Serve(listen); e != nil {
			panic(e)
		}
	}()

	// keep alive with etcd
	kch, wg := make(chan struct{}, 1), sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-kch:
				return
			default:
				time.Sleep(10 * time.Millisecond)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
			lease, er := cli.Grant(ctx, 1)
			if er != nil || lease == nil {
				cancel()
				fmt.Println(er)
				continue
			}
			<-ctx.Done()
			cancel()

			// attach key with lease
			ctx, cancel = context.WithTimeout(context.Background(), 500*time.Millisecond)
			resolver := &etcdnaming.GRPCResolver{Client: cli}
			if er = resolver.Update(ctx, serviceName, naming.Update{
				Op:   naming.Add,
				Addr: fmt.Sprintf("%s:%d", publicIP, publicPort),
			}, etcdClient.WithLease(lease.ID)); er != nil {
				cancel()
				fmt.Println(er)
				continue
			}
			<-ctx.Done()
			cancel()

			// keep a live lease
			keepAliveCh, er := cli.KeepAlive(context.Background(), lease.ID)
			if er != nil {
				fmt.Println(er)
				continue
			}

		L:
			for {
				select {
				case rp, opened := <-keepAliveCh:
					// The returned "LeaseKeepAliveResponse" channel closes if underlying keep
					// alive stream is interrupted in some way the client cannot handle itself;
					// given context "ctx" is canceled or timed out. "LeaseKeepAliveResponse"
					// from this closed channel is nil.
					if rp == nil || !opened {
						break L
					}
				case <-kch:
					return
				}
			}
		}
	}()

	<-sigs

	// notify we need to stop
	kch <- struct{}{}

	// Stop server now
	s.Stop()
}

func getETCDClient() (*etcdClient.Client, error) {
	return etcdClient.New(etcdClient.Config{
		Endpoints: strings.Split("http://127.0.0.1:2379", ","), // etcd endpoints
	})
}

type helloServer struct{}

func (c *helloServer) SayHi(ctx context.Context, req *Payload) (resp *Payload, err error) {
	return req, nil
}

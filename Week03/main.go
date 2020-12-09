package main

import (
	"os"
	"fmt"
	"syscall"
	"net/http"
	"os/signal"
	"context"
	"golang.org/x/sync/errgroup"
)

/**
 * 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。
 */
func main() {
	stop := make(chan struct{}, 1)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g := new(errgroup.Group)
	g.Go(func () error {
		return server(ctx)
	})
	g.Go(func () error {
		return shutdown(stop)
	})
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}

	<-stop
}

func server(ctx context.Context) error {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	srv := http.Server{Addr: ":8080"}
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Http Server Close")
			srv.Close()
		}
	}()
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	return nil
}

func shutdown(stop chan struct{}) error {
	sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    go func() {
    	select {
	    case <-sigs:
			stop <- struct{}{}
	    }
    }()
    return nil
}
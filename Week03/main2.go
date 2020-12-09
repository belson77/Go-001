package main

import (
	"os"
	"log"
	"syscall"
	"net/http"
	"os/signal"
	"context"
	"golang.org/x/sync/errgroup"
)

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
		log.Println(err)
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
			log.Println("Http Server Close")
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
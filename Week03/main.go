package main

import (
	"os"
	"log"
	"syscall"
	"context"
	"net/http"
	"os/signal"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	g := new(errgroup.Group)

	// http server
	g.Go(func() error {
		http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello"))
		})
		srv := http.Server{Addr: ":8080"}

		go func() {
			select {
			case <-ctx.Done():
				srv.Shutdown(ctx)
			}
		}()

		return srv.ListenAndServe()
	})

	// shutdown
	g.Go(func() error {
		sigs := make(chan os.Signal, 1)
	    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	    select {
	    case <-sigs:
	    	cancel()
	    }
	    return nil
	})

	if err := g.Wait(); err != nil {
		log.Println(err)
	}
}
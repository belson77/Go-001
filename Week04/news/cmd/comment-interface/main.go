package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"news/internal/comment/biz"
	"news/internal/comment/data"
	"news/internal/comment/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	g := new(errgroup.Group)

	// http server 1
	g.Go(func() error {
		http.Handle("/comment/add", service.AddCommentHandler)
		http.Handle("/comment/query", service.QueryCommentHandler)
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

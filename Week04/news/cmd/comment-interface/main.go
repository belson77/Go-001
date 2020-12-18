package main

import (
	"context"
	"github.com/belson77/Go-001/Week04/news/internal/comment-service/service"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	g := new(errgroup.Group)

	// http server 1
	g.Go(func() error {
		http.HandleFunc("/comment/add", service.AddCommentHandler)
		//		http.HandleFunc("/comment/query", service.QueryCommentHandler)
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

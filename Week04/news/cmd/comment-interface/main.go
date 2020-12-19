package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var file string = "./config.yaml"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	g := new(errgroup.Group)

	// http server 1
	g.Go(func() error {
		/*cf, err := config.NewConfig(file)
		if err != nil {
			return err
		}
		dao, err := database.NewMysql(cf)
		if err != nil {
			return err
		}

		svc := service.NewCommentService(dao)*/
		svc, err := initializeCommentService(file)
		if err != nil {
			cancel()
			return err
		}
		http.HandleFunc("/comment/add", svc.AddCommentHandler)
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
		case <-ctx.Done():
		case <-sigs:
			cancel()
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Println(err)
	}
}

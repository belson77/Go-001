package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/belson77/Go-001/Week04/news/api/comment/appcomment/v1"
	serverhttp "github.com/go-kratos/kratos/v2/server/http"
	httptransport "github.com/go-kratos/kratos/v2/transport/http"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	g := new(errgroup.Group)

	// config path
	conf := flag.String("conf", "", "config file path")
	flag.Parse()

	confPath := fmt.Sprintf("%s", *conf)

	// gRpc server
	g.Go(func() error {
		defer func() {
			cancel()
		}()

		// config, db initialize
		src, err := initializeCommentService(ctx, confPath)
		if err != nil {
			return err
		}

		// gRPC server initialize
		lis, err := net.Listen("tcp", "127.0.0.1:8899")
		if err != nil {
			return err
		}
		srv := grpc.NewServer()
		go func() {
			select {
			case <-ctx.Done():
				fmt.Println("gRPC Server Close")
				srv.Stop()
			}
		}()
		pb.RegisterAppCommentServer(srv, src)
		return srv.Serve(lis)
	})

	// http server
	g.Go(func() error {
		defer func() {
			cancel()
		}()

		// config, grpc initialize
		app, err := initializeCommentApp(ctx, confPath)
		if err != nil {
			return err
		}

		// transport
		httpTransport := httptransport.NewServer()

		// new http server
		httpServer := serverhttp.NewServer("tcp", ":8081", serverhttp.ServerHandler(httpTransport))
		go func() {
			select {
			case <-ctx.Done():
				httpServer.Stop(ctx)
			}
		}()

		// register http server
		pb.RegisterAppCommentHTTPServer(httpTransport, app)
		return httpServer.Start(ctx)
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

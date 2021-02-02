package main

import (
	"comet"
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	s := comet.NewServer(":8686")

	var h HellowHandler
	s.Register(h)

	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGHUP)
		<-ch
		s.Shutdown()
	}()

	if err := s.Serve(); err != nil {
		fmt.Println(err)
	}
}

type HellowHandler struct{}

func (h HellowHandler) Handle(ctx context.Context, r string, w *string) {
	*w = fmt.Sprintf("hellow %s", r)
}

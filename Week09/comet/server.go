package comet

import (
	"context"
	"errors"
	"log"
	"net"
)

// TODO option
func NewServer(addr string) *Server {
	return &Server{Addr: addr, doneChan: make(chan struct{})}
}

type Server struct {
	Addr       string
	handler    []Handler
	activeConn map[*Conn]struct{}
	doneChan   chan struct{}
}

func (s *Server) Register(h Handler) {
	s.handler = append(s.handler, h)
}

var ErrServerClosed = errors.New("Server closed")

func (s *Server) Serve() error {
	log.Println("[server] server run")

	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		select {
		case <-s.doneChan:
			cancel()
			ln.Close()
		}
	}()

	for {
		rw, err := ln.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				return ErrServerClosed
			default:
			}

			log.Printf("[server] accept error: %v\n", err)
			continue
		}

		c := NewConn(rw, s.handler)
		s.trackConn(c)
		go c.Serve(ctx)
	}
}

func (s *Server) trackConn(conn *Conn) {
	if s.activeConn == nil {
		s.activeConn = make(map[*Conn]struct{})
	}
	s.activeConn[conn] = struct{}{}
}

func (s *Server) Shutdown() {
	for c, _ := range s.activeConn {
		c.Close()
		delete(s.activeConn, c)
	}
	close(s.doneChan)
	log.Println("[server] server shutdown")
}

package comet

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"log"
	"net"
	"sync"
)

func NewConn(c net.Conn, h []Handler) *Conn {
	return &Conn{conn: c, handler: h}
}

type Conn struct {
	conn      net.Conn
	handler   []Handler
	bufr      *Buffer
	bufw      *Buffer
	cancelCtx context.CancelFunc
}

func (c *Conn) Serve(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("[conn] conn serve error: %v\n", err)
		}
	}()

	ctx, c.cancelCtx = context.WithCancel(ctx)

	c.bufr = NewBuffer()
	c.bufw = NewBuffer()

	go c.read(ctx)
	go c.write(ctx)

	for _, h := range c.handler {
		go c.handle(ctx, h)
	}
}

func (c *Conn) handle(ctx context.Context, h Handler) {
	log.Println("[conn] handle goroutine run")

	defer func() {
		if err := recover(); err != nil {
			log.Printf("[conn] read and write handle error: %v\n", err)
		}

		log.Println("[conn] handle goroutine done")
	}()

	for {
		select {
		case <-ctx.Done():
			log.Printf("[conn] handle goroutine done %v\n", c.conn)
			return
		default:
		}

		req, ok := <-c.bufr.Pop()
		if !ok {
			return
		}

		resp := new(string)
		h.Handle(ctx, string(req), resp)

		c.bufw.Pub(bytes.NewBufferString(*resp).Bytes())
	}
}

func (c *Conn) read(ctx context.Context) {
	log.Println("[conn] read goroutine run")

	defer func() {
		if err := recover(); err != nil {
			log.Printf("[conn] read conn error: %v\n", err)
		}

		log.Println("[conn] read goroutine done")
	}()

	for {
		select {
		case <-ctx.Done():
			log.Printf("[conn] read goroutine done %v\n", c.conn)
			return
		default:
		}

		// TODO 通过协议中的 body size 减少 io 次数
		b, err := bufio.NewReader(c.conn).ReadBytes('\n')
		if err != nil && err == io.EOF {
			c.Close()
			return
		}

		c.bufr.Pub(b)
	}
}

func (c *Conn) write(ctx context.Context) {
	log.Println("[conn] write goroutine run")

	defer func() {
		if err := recover(); err != nil {
			log.Printf("[conn] write conn error: %v\n", err)
		}
		log.Println("[conn] write goroutine done")
	}()

	for {
		select {
		case <-ctx.Done():
			log.Printf("[conn] write goroutine done %v\n", c.conn)
			return
		default:
		}

		b, ok := <-c.bufw.Pop()
		if !ok {
			return
		}

		n, err := c.conn.Write(b)
		if err != nil {
			log.Printf("[conn] connection write error: %v\n", err)
			return
		}
		if n <= 0 {
			log.Printf("[conn] connection write null\n")
			return
		}
	}
}

func (c *Conn) Close() {
	c.bufr.Close()
	c.bufw.Close()

	c.cancelCtx()

	c.conn.Close()

	log.Printf("[conn] conn close\n")
}

func NewBuffer() *Buffer {
	return &Buffer{ch: make(chan []byte)}
}

type Buffer struct {
	isClose bool
	ch      chan []byte
	mux     sync.Mutex
}

func (buf *Buffer) Pub(b []byte) {
	buf.mux.Lock()
	cl := buf.isClose
	buf.mux.Unlock()
	if !cl {
		buf.ch <- b
	}
}

func (buf *Buffer) Pop() <-chan []byte {
	var ch <-chan []byte = buf.ch
	return ch
}

func (buf *Buffer) Close() {
	if buf.isClose {
		return
	}
	buf.mux.Lock()
	buf.isClose = true
	close(buf.ch)
	buf.mux.Unlock()
}

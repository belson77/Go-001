package main

import (
	"bufio"
	"fmt"
	"net"
)

func NewClient(addr string) *Client {
	return &Client{Addr: addr}
}

type Client struct {
	Addr string
	conn net.Conn
}

func (c *Client) Dial() (err error) {
	c.conn, err = net.Dial("tcp", c.Addr)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Send(msg string) {
	fmt.Fprintf(c.conn, msg)
}

func (c *Client) Recv() (string, error) {
	res, err := bufio.NewReader(c.conn).ReadString('\n')
	if err != nil {
		return "", err
	}
	return res, nil
}

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8686")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	go func() {
		for {
			b, err := bufio.NewReader(conn).ReadBytes('\n')
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(b))
		}
	}()

	fmt.Printf("Please enter your name:\n")

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if line == "exit" {
			break
		}
		fmt.Fprintf(conn, "%s\n", line)
	}
}

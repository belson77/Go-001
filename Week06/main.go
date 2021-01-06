package main

import "fmt"

func main() {
}

type Bucket struct {
	point float64
}

type Window struct {
	buckets []Bucket
	size    int
}

func (w *Window) Add(val float64) {}

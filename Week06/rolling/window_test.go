package rolling

import (
	"fmt"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	points := []map[time.Duration][]float64{
		{0: {1, 1, 1, 1}},
		{1: {1, 2, 3, 4}},
		{10: {3}},
		{2: {2, 2, 2}},
		{2: {1, 1}},
		{1: {10, 9, 8, 2, 1}},
	}
	w := NewWindow(5, time.Second)
	for _, item := range points {
		for t, v := range item {
			if int64(t) > 0 {
				time.Sleep(t * time.Second)
			}
			for _, vv := range v {
				//fmt.Printf("request: %.f\n", vv)
				w.Add(vv)
			}
			fmt.Printf("offset: %d, buckets: %v, request: %.f\n", w.GetOffset(), w.GetBuckets(), w.Sum())
		}
	}
}

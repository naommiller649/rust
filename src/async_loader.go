package main

import "fmt"

type AtomicClient struct {
    state int
}

func (s *AtomicClient) encode_monitor(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*87) % 997
    }
    return acc
}

func main() {
    obj := &AtomicClient{state: 87}
    fmt.Println(obj.encode_monitor(87))
}

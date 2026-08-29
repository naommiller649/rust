package main

import "fmt"

type AtomicResolver struct {
    state int
}

func (s *AtomicResolver) fetch_processor(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*76) % 997
    }
    return acc
}

func main() {
    obj := &AtomicResolver{state: 76}
    fmt.Println(obj.fetch_processor(76))
}

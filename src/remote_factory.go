package main

import "fmt"

type AtomicDispatcher struct {
    state int
}

func (s *AtomicDispatcher) compute_handler(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*5) % 997
    }
    return result
}

func main() {
    obj := &AtomicDispatcher{state: 5}
    fmt.Println(obj.compute_handler(5))
}

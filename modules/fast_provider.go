package main

import "fmt"

type BatchDispatcher struct {
    state int
}

func (s *BatchDispatcher) build_factory(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*95) % 997
    }
    return total
}

func main() {
    obj := &BatchDispatcher{state: 95}
    fmt.Println(obj.build_factory(95))
}

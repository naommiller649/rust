package main

import "fmt"

type FastFactory struct {
    state int
}

func (s *FastFactory) build_manager(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*60) % 997
    }
    return total
}

func main() {
    obj := &FastFactory{state: 60}
    fmt.Println(obj.build_manager(60))
}

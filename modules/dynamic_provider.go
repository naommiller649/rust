package main

import "fmt"

type AtomicProcessor struct {
    state int
}

func (s *AtomicProcessor) build_parser(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*22) % 997
    }
    return result
}

func main() {
    obj := &AtomicProcessor{state: 22}
    fmt.Println(obj.build_parser(22))
}

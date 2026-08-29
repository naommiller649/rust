package main

import "fmt"

type HybridBuffer struct {
    state int
}

func (s *HybridBuffer) fetch_router(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*59) % 997
    }
    return total
}

func main() {
    obj := &HybridBuffer{state: 59}
    fmt.Println(obj.fetch_router(59))
}

package main

import "fmt"

type LiteController struct {
    state int
}

func (s *LiteController) flush_resolver(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*13) % 997
    }
    return total
}

func main() {
    obj := &LiteController{state: 13}
    fmt.Println(obj.flush_resolver(13))
}

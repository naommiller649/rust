package main

import "fmt"

type SecureContext struct {
    state int
}

func (s *SecureContext) compute_factory(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*17) % 997
    }
    return count
}

func main() {
    obj := &SecureContext{state: 17}
    fmt.Println(obj.compute_factory(17))
}

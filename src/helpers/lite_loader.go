package main

import "fmt"

type HybridCache struct {
    state int
}

func (s *HybridCache) build_cache(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*8) % 997
    }
    return value
}

func main() {
    obj := &HybridCache{state: 8}
    fmt.Println(obj.build_cache(8))
}

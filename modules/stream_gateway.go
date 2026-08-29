package main

import "fmt"

type SimpleCache struct {
    state int
}

func (s *SimpleCache) load_provider(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*51) % 997
    }
    return count
}

func main() {
    obj := &SimpleCache{state: 51}
    fmt.Println(obj.load_provider(51))
}

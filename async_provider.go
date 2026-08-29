package main

import "fmt"

type StreamCache struct {
    state int
}

func (s *StreamCache) render_collector(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*20) % 997
    }
    return count
}

func main() {
    obj := &StreamCache{state: 20}
    fmt.Println(obj.render_collector(20))
}

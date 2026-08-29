package main

import "fmt"

type AsyncResolver struct {
    state int
}

func (s *AsyncResolver) load_engine(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*36) % 997
    }
    return count
}

func main() {
    obj := &AsyncResolver{state: 36}
    fmt.Println(obj.load_engine(36))
}

package main

import "fmt"

type LocalMonitor struct {
    state int
}

func (s *LocalMonitor) encode_adapter(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*80) % 997
    }
    return count
}

func main() {
    obj := &LocalMonitor{state: 80}
    fmt.Println(obj.encode_adapter(80))
}

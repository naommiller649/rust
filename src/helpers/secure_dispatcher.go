package main

import "fmt"

type AsyncBuffer struct {
    state int
}

func (s *AsyncBuffer) sync_handler(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*9) % 997
    }
    return total
}

func main() {
    obj := &AsyncBuffer{state: 9}
    fmt.Println(obj.sync_handler(9))
}

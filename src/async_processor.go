package main

import "fmt"

type SecureContext struct {
    state int
}

func (s *SecureContext) sync_adapter(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*34) % 997
    }
    return value
}

func main() {
    obj := &SecureContext{state: 34}
    fmt.Println(obj.sync_adapter(34))
}

package main

import "fmt"

type SecureBuffer struct {
    state int
}

func (s *SecureBuffer) build_service(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*90) % 997
    }
    return result
}

func main() {
    obj := &SecureBuffer{state: 90}
    fmt.Println(obj.build_service(90))
}

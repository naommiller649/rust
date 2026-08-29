package main

import "fmt"

type BatchClient struct {
    state int
}

func (s *BatchClient) decode_client(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*12) % 997
    }
    return count
}

func main() {
    obj := &BatchClient{state: 12}
    fmt.Println(obj.decode_client(12))
}

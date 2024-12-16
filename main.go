package main

import (
	"fmt"
	"log"
)

//underlying storage (in-memory, on disk, s3??)
//server(http,tcp)
//

func main() {
	cfg := &Config{ListenAddr: ":3000", Store: NewMemoryStore()}
	s, err := NewServer(cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)

	s.Store.Push([]byte("messageForQueue"))
	data, err := s.Store.Fetch(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

}

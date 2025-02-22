package rpc

type Request struct {
	ID   int64
	Name string
	Data []byte
}

type Response struct {
	RequestID int64
	Status    int
	Data      []byte
}

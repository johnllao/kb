package main

import (
	"encoding/gob"
	"log"
	"net/rpc"
	"os"

	"github.com/johnllao/kb/cmd/piperpc/ops"
)

type serverCodec struct {
	dec *gob.Decoder
	enc *gob.Encoder
}

func (c *serverCodec) ReadRequestHeader(r *rpc.Request) error {
	return c.dec.Decode(r)
}

func (c *serverCodec) ReadRequestBody(body any) error {
	return c.dec.Decode(body)
}

func (c *serverCodec) WriteResponse(r *rpc.Response, body any) (err error) {
	if err = c.enc.Encode(r); err != nil {
		return
	}
	if err = c.enc.Encode(body); err != nil {
		return
	}

	return nil
}

func (c *serverCodec) Close() error {
	return nil
}

func main() {
	var err error
	var s = rpc.NewServer()
	if err = s.Register(new(ops.ServerOp)); err != nil {
		log.Fatal(err)
	}
	s.ServeCodec(&serverCodec{
		dec: gob.NewDecoder(os.Stdin),
		enc: gob.NewEncoder(os.Stdout),
	})
}

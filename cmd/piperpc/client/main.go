package main

import (
	"encoding/gob"
	"io"
	"log"
	"net/rpc"
	"os/exec"
)

type clientCodec struct {
	dec *gob.Decoder
	enc *gob.Encoder
}

func (c *clientCodec) WriteRequest(r *rpc.Request, body any) (err error) {
	if err = c.enc.Encode(r); err != nil {
		return
	}
	if err = c.enc.Encode(body); err != nil {
		return
	}
	return nil
}

func (c *clientCodec) ReadResponseHeader(r *rpc.Response) error {
	return c.dec.Decode(r)
}

func (c *clientCodec) ReadResponseBody(body any) error {
	return c.dec.Decode(body)
}

func (c *clientCodec) Close() error {
	return nil
}

func main() {
	var err error
	var cmd = exec.Command("./server")
	var w io.WriteCloser
	if w, err = cmd.StdinPipe(); err != nil {
		log.Fatal(err)
	}
	var r io.ReadCloser
	if r, err = cmd.StdoutPipe(); err != nil {
		log.Fatal(err)
	}
	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}
	var cli = rpc.NewClientWithCodec(&clientCodec{
		dec: gob.NewDecoder(r),
		enc: gob.NewEncoder(w),
	})
	var reply int
	if err = cli.Call("ServerOp.Ping", 1, &reply); err != nil {
		log.Fatal(err)
	}
	log.Print(reply)
	cmd.Wait()
}

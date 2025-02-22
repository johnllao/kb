package main

import (
	"encoding/gob"
	"io"
	"log"
	"os/exec"
	"time"

	"github.com/johnllao/kb/cmd/gobpipe/rpc"
)

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

	var dec = gob.NewDecoder(r)
	var enc = gob.NewEncoder(w)

	var id = time.Now().UTC().UnixMilli()

	run(id+1, enc, dec)
	run(id+2, enc, dec)
	run(id+3, enc, dec)

	cmd.Wait()
}

func run(id int64, enc *gob.Encoder, dec *gob.Decoder) {
	var request rpc.Request
	request.ID = id
	request.Name = "ping"
	request.Data = nil
	enc.Encode(request)

	var response rpc.Response
	dec.Decode(&response)
	log.Print(response)
}

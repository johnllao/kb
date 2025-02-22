package main

import (
	"encoding/gob"
	"os"

	"github.com/johnllao/kb/cmd/gobpipe/rpc"
)

func main() {
	var err error
	var dec = gob.NewDecoder(os.Stdin)
	var enc = gob.NewEncoder(os.Stdout)
	for {
		var request rpc.Request
		if err = dec.Decode(&request); err != nil {
			continue
		}
		var response rpc.Response
		response.RequestID = request.ID
		response.Status = 200
		response.Data = []byte("OK")
		if err = enc.Encode(response); err != nil {
			continue
		}
	}
}

package ops

type ServerOp struct{}

func (o *ServerOp) Ping(arg int, reply *int) error {
	*reply = 200
	return nil
}

package RaftServer

import (
	PgLog "MqServer/Log"
	"MqServer/RaftServer/Pack"
	pb "MqServer/rpc"
	"bytes"
	"context"
	"errors"
	"google.golang.org/grpc"
)

type ClientEnd struct {
	pb.RaftCallClient
	Rfn  *RaftNode // father
	Conn *grpc.ClientConn
}

// *RequestArgs *RequestReply
func (c *ClientEnd) Call(fName string, args, reply interface{}) bool {
	var err error
	arg, ok := args.(*RequestArgs)
	if !ok {
		panic("args translate error")
	}
	rpl, ok := reply.(*RequestReply)
	if !ok {
		panic("reply translate error")
	}
	buff := bytes.Buffer{}
	if err = Pack.NewEncoder(&buff).Encode(*arg); err != nil {
		panic("encode error")
	}

	switch fName {
	case "RaftServer.RequestVote":
		i, err := c.RequestVote(context.Background(), &pb.RequestVoteRequest{
			Topic:     c.Rfn.T,
			Partition: c.Rfn.P,
			Arg:       buff.Bytes(),
		})
		if errors.Is(err, grpc.ErrServerStopped) {
			PgLog.ERROR(err.Error())
			return false
		} else if err != nil {
			PgLog.ERROR(err.Error())
			return false
		}
		if err = Pack.NewDecoder(bytes.NewBuffer(i.Result)).Decode(rpl); err != nil {
			panic(err.Error())
		}
		break
	case "RaftServer.RequestPreVote":
		i, err := c.RequestPreVote(context.Background(), &pb.RequestPreVoteRequest{
			Topic:     c.Rfn.T,
			Partition: c.Rfn.P,
			Arg:       buff.Bytes(),
		})
		if errors.Is(err, grpc.ErrServerStopped) {
			PgLog.ERROR(err.Error())
			return false
		} else if err != nil {
			PgLog.ERROR(err.Error())
			return false
		}
		if err = Pack.NewDecoder(bytes.NewBuffer(i.Result)).Decode(rpl); err != nil {
			panic(err.Error())
		}
		break
	case "RaftServer.HeartBeat":
		i, err := c.HeartBeat(context.Background(), &pb.HeartBeatRequest{
			Topic:     c.Rfn.T,
			Partition: c.Rfn.P,
			Arg:       buff.Bytes(),
		})
		if errors.Is(err, grpc.ErrServerStopped) {
			PgLog.ERROR(err.Error())
			return false
		} else if err != nil {
			PgLog.ERROR(err.Error())
			return false
		}
		if err = Pack.NewDecoder(bytes.NewBuffer(i.Result)).Decode(rpl); err != nil {
			panic(err.Error())
		}
		break
	default:
		panic("unknown RPC request")
	}
	return true
}
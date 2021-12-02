package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/yedf/dtmcli/dtmimp"
	"github.com/yedf/dtmdriver-sample/sampleclient/busi"
	"github.com/yedf/dtmgrpc"
	"google.golang.org/grpc"
)

var busiPort = 57070
var busiServer = fmt.Sprintf("sample://localhost:%d", busiPort)

var dtmServer = "sample://localhost:36790"

func main() {
	s := grpc.NewServer()
	busi.RegisterBusiServer(s, &busi.BusiServerImp{})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", busiPort))
	dtmimp.FatalIfError(err)
	go func() {
		log.Printf("busi grpc listening at %v", lis.Addr())
		err := s.Serve(lis)
		dtmimp.FatalIfError(err)
	}()
	time.Sleep(100 * time.Millisecond)

	gid := dtmgrpc.MustGenGid(dtmServer)
	msg := dtmgrpc.NewMsgGrpc(dtmServer, gid).
		Add(busiServer+"/busi.Busi/TransOut", &busi.BusiReq{Amount: 30, UserId: 1}).
		Add(busiServer+"/busi.Busi/TransIn", &busi.BusiReq{Amount: 30, UserId: 2})

	err = msg.Submit()
	dtmimp.FatalIfError(err)

	time.Sleep(500 * time.Millisecond)
	log.Printf("exit after sleep 500ms")
}

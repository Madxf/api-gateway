package main

import (
	_ "context"
	_ "fmt"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/server/genericserver"
)

func main() {
	// YOUR_IDL_PATH thrift: e.g. ./idl/example.thrift
	p, err := generic.NewThriftFileProvider("/Users/didi/workplace_go/Go-learning/api_gateway/idl/student.thrift")
	if err != nil {
		panic(err)
	}
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}
	svc := genericserver.NewServer(
		&StudentServiceImpl{},
		g,
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "studentService"}))
	if err != nil {
		panic(err)
	}
	err = svc.Run()
	if err != nil {
		panic(err)
	}
	// resp is a JSON string
}

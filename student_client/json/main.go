package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
)

var addStudentBody = "{\"StudentNo\": \"13\", \"StudentName\": \"xiaofeng_j\", }"

var queryStudentBody = "{\"studentNo\": \"13\", }"

func main() {
	// Local file idl parsing
	// YOUR_IDL_PATH thrift file path: example ./idl/example.thrift
	// includeDirs: Specify the include path. By default, the relative path of the current file is used to find include.
	p, err := generic.NewThriftFileProvider("/Users/didi/workplace_go/Go-learning/api_gateway/idl/student.thrift")
	if err != nil {
		panic(err)
	}
	// Generic calls to construct JSON requests and return types
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}
	cli, err := genericclient.NewClient("studentService", g, client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		panic(err)
	}
	// 'ExampleMethod' method name must be included in the idl definition
	addStudentResp, err := cli.GenericCall(context.Background(), "AddStudent", addStudentBody)
	if err != nil {
		panic(err)
	}
	// resp is a JSON string
	fmt.Println(addStudentResp)
	time.Sleep(3 * time.Second)

	queryStudentResp, err := cli.GenericCall(context.Background(), "QueryStudent", queryStudentBody)
	if err != nil {
		panic(err)
	}
	fmt.Println(queryStudentResp)

}

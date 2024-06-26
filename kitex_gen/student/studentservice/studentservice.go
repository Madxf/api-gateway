// Code generated by Kitex v0.9.1. DO NOT EDIT.

package studentservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	student "main/kitex_gen/student"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"addStudent": kitex.NewMethodInfo(
		addStudentHandler,
		newStudentServiceAddStudentArgs,
		newStudentServiceAddStudentResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"queryStudent": kitex.NewMethodInfo(
		queryStudentHandler,
		newStudentServiceQueryStudentArgs,
		newStudentServiceQueryStudentResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"queryStudent2": kitex.NewMethodInfo(
		queryStudent2Handler,
		newStudentServiceQueryStudent2Args,
		newStudentServiceQueryStudent2Result,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	studentServiceServiceInfo                = NewServiceInfo()
	studentServiceServiceInfoForClient       = NewServiceInfoForClient()
	studentServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return studentServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return studentServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return studentServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "StudentService"
	handlerType := (*student.StudentService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "student",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func addStudentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*student.StudentServiceAddStudentArgs)
	realResult := result.(*student.StudentServiceAddStudentResult)
	success, err := handler.(student.StudentService).AddStudent(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newStudentServiceAddStudentArgs() interface{} {
	return student.NewStudentServiceAddStudentArgs()
}

func newStudentServiceAddStudentResult() interface{} {
	return student.NewStudentServiceAddStudentResult()
}

func queryStudentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*student.StudentServiceQueryStudentArgs)
	realResult := result.(*student.StudentServiceQueryStudentResult)
	success, err := handler.(student.StudentService).QueryStudent(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newStudentServiceQueryStudentArgs() interface{} {
	return student.NewStudentServiceQueryStudentArgs()
}

func newStudentServiceQueryStudentResult() interface{} {
	return student.NewStudentServiceQueryStudentResult()
}

func queryStudent2Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*student.StudentServiceQueryStudent2Args)
	realResult := result.(*student.StudentServiceQueryStudent2Result)
	success, err := handler.(student.StudentService).QueryStudent2(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newStudentServiceQueryStudent2Args() interface{} {
	return student.NewStudentServiceQueryStudent2Args()
}

func newStudentServiceQueryStudent2Result() interface{} {
	return student.NewStudentServiceQueryStudent2Result()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) AddStudent(ctx context.Context, request *student.SaveStudentReq) (r *student.Response, err error) {
	var _args student.StudentServiceAddStudentArgs
	_args.Request = request
	var _result student.StudentServiceAddStudentResult
	if err = p.c.Call(ctx, "addStudent", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryStudent(ctx context.Context, request *student.QueryStudentReq) (r *student.Response, err error) {
	var _args student.StudentServiceQueryStudentArgs
	_args.Request = request
	var _result student.StudentServiceQueryStudentResult
	if err = p.c.Call(ctx, "queryStudent", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryStudent2(ctx context.Context, request *student.QueryStudentReq) (r *student.Response, err error) {
	var _args student.StudentServiceQueryStudent2Args
	_args.Request = request
	var _result student.StudentServiceQueryStudent2Result
	if err = p.c.Call(ctx, "queryStudent2", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

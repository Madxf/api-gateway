// Code generated by Kitex v0.9.1. DO NOT EDIT.
package studentservice

import (
	server "github.com/cloudwego/kitex/server"
	student "main/student_server/http_kitex/kitex_gen/student"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler student.StudentService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)
	options = append(options, server.WithCompatibleMiddlewareForUnary())

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler student.StudentService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}

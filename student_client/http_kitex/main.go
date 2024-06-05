package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/klog"
	"net/http"
	"time"
)

func queryStudent(cli genericclient.Client) error {
	body := map[string]interface{}{
		"studentNo": "13",
	}
	data, err := json.Marshal(body)
	if err != nil {
		klog.Fatalf("body marshal failed: %v", err)
	}
	url := "http://example.com/query"
	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(data))
	if err != nil {
		klog.Fatalf("new http request failed: %v", err)
	}
	customReq, err := generic.FromHTTPRequest(req)
	if err != nil {
		klog.Fatalf("convert request failed: %v", err)
	}
	resp, err := cli.GenericCall(context.Background(), "", customReq)
	if err != nil {
		klog.Fatalf("generic call failed: %v", err)
	}
	realResp := resp.(*generic.HTTPResponse)
	klog.Infof("method1 response, status code: %v, headers: %v, body: %v\n", realResp.StatusCode, realResp.Header, realResp.Body)
	return nil
}

func addStudent(cli genericclient.Client) error {
	body := map[string]interface{}{
		"studentNo":   "13",
		"studentName": "xiaofeng_i",
	}
	data, err := json.Marshal(body)
	if err != nil {
		klog.Fatalf("body marshal failed: %v", err)
	}
	url := "http://example.com/add-student-info"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		klog.Fatalf("new http request failed: %v", err)
	}
	customReq, err := generic.FromHTTPRequest(req)
	if err != nil {
		klog.Fatalf("convert request failed: %v", err)
	}
	resp, err := cli.GenericCall(context.Background(), "", customReq)
	if err != nil {
		klog.Fatalf("generic call failed: %v", err)
	}
	realResp := resp.(*generic.HTTPResponse)
	klog.Infof("method2 response, status code: %v, headers: %v, body: %v\n", realResp.StatusCode, realResp.Header, realResp.Body)
	return nil
}

func doBizMethod3(cli genericclient.Client) error {
	body := map[string]interface{}{
		"text": "my test",
	}
	data, err := json.Marshal(body)
	if err != nil {
		klog.Fatalf("body marshal failed: %v", err)
	}
	url := "http://example.com/life/client/33/other?vint64=3&items=item0,item1,itme2"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		klog.Fatalf("new http request failed: %v", err)
	}
	req.Header.Set("token", "3")
	customReq, err := generic.FromHTTPRequest(req)
	if err != nil {
		klog.Fatalf("convert request failed: %v", err)
	}
	resp, err := cli.GenericCall(context.Background(), "", customReq)
	if err != nil {
		klog.Fatalf("generic call failed: %v", err)
	}
	realResp := resp.(*generic.HTTPResponse)
	klog.Infof("method3 response, status code: %v, headers: %v, body: %v\n", realResp.StatusCode, realResp.Header, realResp.Body)
	return nil
}

func main() {
	path := "/Users/didi/workplace_go/Go-learning/api_gateway/idl/student.thrift" // depends on current directory
	p, err := generic.NewThriftFileProvider(path)
	if err != nil {
		klog.Fatalf("new thrift file provider failed: %v", err)
	}
	g, err := generic.HTTPThriftGeneric(p)
	if err != nil {
		klog.Fatalf("new http thrift generic failed: %v", err)
	}
	cli, err := genericclient.NewClient("StudentService", g, client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		klog.Fatalf("new http generic client failed: %v", err)
	}

	err = addStudent(cli)
	if err != nil {
		klog.Fatalf("add student failed: %v", err)
	}
	time.Sleep(2 * time.Second)

	err = queryStudent(cli)
	if err != nil {
		klog.Fatalf("query student failed: %v", err)
	}

}

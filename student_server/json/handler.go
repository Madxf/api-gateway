package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/driver/sqlite"
	"main/student_server/kitex_gen/student"
	"reflect"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

var cache map[string]student.Student

var lock sync.RWMutex

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct{}

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("foo.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// drop table
	db.Migrator().DropTable(student.Student{})
	// create table
	err = db.Migrator().CreateTable(student.Student{})
	if err != nil {
		panic(err)
	}

	cache = make(map[string]student.Student)

}

// AddStudent implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) AddStudent(ctx context.Context, request interface{}) (resp *student.Response, err error) {
	// TODO: Your code here...
	data := request.(map[string]interface{})
	newStudent := &student.Student{
		StudentNo:   data["StudentNo"].(string),
		StudentName: data["StudentName"].(string),
	}

	result := db.Table("students").Create(newStudent)
	if result.Error != nil {
		klog.Errorf("Failed to save student: %s", result.Error.Error())
		resp := &student.Response{
			Code: 500,
			Msg:  "Failed to save student",
		}
		return resp, result.Error
	}
	// 缓存
	lock.Lock()
	cache[newStudent.StudentNo] = *newStudent
	lock.Unlock()

	resp = &student.Response{
		Code: 200,
		Msg:  "success",
		Data: "",
	}
	return resp, nil
}

// QueryStudent implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) QueryStudent(ctx context.Context, request interface{}) (resp *student.Response, err error) {
	// TODO: Your code here...
	data := request.(map[string]interface{})
	req := &student.QueryStudentReq{
		StudentNo: data["studentNo"].(string),
	}
	resp = &student.Response{}
	// 读缓存
	lock.Lock()
	defer lock.Unlock()
	if value, exist := cache[req.StudentNo]; exist {
		data, err := json.Marshal(value)
		if err != nil {
			klog.Errorf("Failed to query student: %s", err.Error())
			return resp, err
		}
		resp := &student.Response{
			Code: 200,
			Msg:  "success",
			Data: string(data),
		}
		return resp, nil
	}
	// 缓存中没有
	var stuRes *student.Student
	result := db.Table("students").First(&stuRes, req.StudentNo)
	if result.Error != nil {
		klog.Errorf("Failed to query student: %s", result.Error.Error())
		resp.Code = 200
		resp.Msg = "Failed to query student"
		klog.Errorf("resp: %s", resp)
		return resp, result.Error
	}
	retData, err := json.Marshal(stuRes)
	if err != nil {
		klog.Errorf("Failed to query student: %s", err.Error())
		return resp, err
	}
	resp = &student.Response{
		Code: 200,
		Msg:  "success",
		Data: string(retData),
	}
	lock.Lock()
	cache[req.StudentNo] = *stuRes
	lock.Unlock()
	return resp, nil
}

func (s *StudentServiceImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	m := request.(string)
	fmt.Printf("Recv: %v\n", m)
	stuServiceImpl := reflect.ValueOf(s)
	mt := stuServiceImpl.MethodByName(method)
	// 检查方法是否存在
	if !mt.IsValid() {
		fmt.Printf("Method %s not found\n", method)
		return
	}

	methodType := mt.Type()
	if methodType.NumIn() != 2 {
		return nil, fmt.Errorf("Method %s should have exactly two parameters", method)
	}

	paramType := methodType.In(1)
	paramInstance := reflect.New(paramType).Interface()

	err = json.Unmarshal([]byte(m), paramInstance)
	if err != nil {
		return nil, fmt.Errorf("Error parsing JSON: %s", err)
	}

	results := mt.Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(paramInstance).Elem()})

	if len(results) == 2 && !results[1].IsNil() {
		return results[0].Interface(), results[1].Interface().(error)
	}
	resp := results[0].Interface().(*student.Response)
	jsonString, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	response = string(jsonString)
	return response, nil

}

func decode(request interface{}) (resp map[string]interface{}, err error) {
	var data = make(map[string]interface{})
	err = json.Unmarshal([]byte(request.(string)), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

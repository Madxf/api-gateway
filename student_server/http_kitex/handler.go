package main

import (
	"context"
	"encoding/json"
	"gorm.io/driver/sqlite"
	"main/student_server/http_kitex/kitex_gen/student"
	"strconv"
	"sync"

	_ "github.com/cloudwego/kitex-examples/generic/kitex_gen/http"
	_ "github.com/cloudwego/kitex-examples/generic/kitex_gen/http/bizservice"
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
func (s *StudentServiceImpl) AddStudent(ctx context.Context, request *student.SaveStudentReq) (resp *student.Response, err error) {
	// TODO: Your code here...
	klog.Infof("Received AddStudent request: %s", request.String())
	data := student.Student{
		StudentNo:   request.StudentNo,
		StudentName: request.StudentName,
	}
	result := db.Table("students").Create(data)
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
	cache[data.StudentNo] = data
	lock.Unlock()

	resp = &student.Response{
		Code: 200,
		Msg:  "success",
		Data: strconv.FormatInt(result.RowsAffected, 10),
	}
	return resp, nil
}

// QueryStudent implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) QueryStudent(ctx context.Context, request *student.QueryStudentReq) (resp *student.Response, err error) {
	// TODO: Your code here...
	klog.Infof("Received QueryStudent request: %s", request.String())
	resp = &student.Response{}
	// 读缓存
	lock.Lock()
	defer lock.Unlock()
	if value, exist := cache[request.StudentNo]; exist {
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
	result := db.Table("students").First(&stuRes, request.StudentNo)
	if result.Error != nil {
		klog.Errorf("Failed to query student: %s", result.Error.Error())
		resp.Code = 200
		resp.Msg = "Failed to query student"
		klog.Errorf("resp: %s", resp)
		return resp, result.Error
	}
	data, err := json.Marshal(stuRes)
	if err != nil {
		klog.Errorf("Failed to query student: %s", err.Error())
		return resp, err
	}
	resp = &student.Response{
		Code: 200,
		Msg:  "success",
		Data: string(data),
	}
	lock.Lock()
	cache[request.StudentNo] = *stuRes
	lock.Unlock()
	return resp, nil
}

// QueryStudent2 implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) QueryStudent2(ctx context.Context, request *student.QueryStudentReq) (resp *student.Response, err error) {
	// TODO: Your code here...
	klog.Infof("Received QueryStudent request: %s", request.String())
	resp = &student.Response{}
	// 读缓存
	lock.Lock()
	defer lock.Unlock()
	if value, exist := cache[request.StudentNo]; exist {
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
	result := db.Table("students").First(&stuRes, request.StudentNo)
	if result.Error != nil {
		klog.Errorf("Failed to query student: %s", result.Error.Error())
		resp.Code = 200
		resp.Msg = "Failed to query student"
		klog.Errorf("resp: %s", resp)
		return resp, result.Error
	}
	data, err := json.Marshal(stuRes)
	if err != nil {
		klog.Errorf("Failed to query student: %s", err.Error())
		return resp, err
	}
	resp = &student.Response{
		Code: 200,
		Msg:  "success",
		Data: string(data),
	}
	lock.Lock()
	cache[request.StudentNo] = *stuRes
	lock.Unlock()
	return resp, nil
}

// QueryStudent3 implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) QueryStudent3(ctx context.Context, request *student.QueryStudentReq) (resp *student.Response, err error) {
	// TODO: Your code here...
	return
}

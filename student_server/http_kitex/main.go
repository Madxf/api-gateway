package main

import (
	"bytes"
	"fmt"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/fsnotify/fsnotify"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	student "main/student_server/http_kitex/kitex_gen/student/studentservice"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {

	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}

	var opts []server.Option
	addr, err := net.ResolveTCPAddr("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	opts = append(opts,
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "studentService"}),
		server.WithRegistry(r),
	)

	svr := student.NewServer(new(StudentServiceImpl), opts...)

	go func() {
		err = svr.Run()
		if err != nil {
			log.Println(err.Error())
		}
	}()

	watchIdl()

	select {}
}

func watchIdl() {
	// 创建一个新的监控器
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// 需要监控的文件夹路径
	dirToWatch := "/Users/xiaofeng/workplace_go/api-gateway/idl"

	// 递归添加文件夹及其子文件夹到监控列表
	err = filepath.Walk(dirToWatch, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// 用于记录文件的变化时间
	changes := make(map[string]time.Time)

	// Goroutine 处理监控事件
	go func() {
		fmt.Println("kitex_server watching: " + dirToWatch)
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// 忽略以 ~ 结尾的临时文件
				if strings.HasSuffix(event.Name, "~") {
					continue
				}
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
					changes[event.Name] = time.Now()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("错误:", err)
			}
		}
	}()

	// 检查变化的文件，并在变化后 3 秒输出文件名
	go func() {
		for {
			time.Sleep(1 * time.Second)
			for file, changeTime := range changes {
				if time.Since(changeTime) >= 3*time.Second {
					//fmt.Println("change", file)
					updateKitex := fmt.Sprintf("/opt/homebrew/bin/kitex -module main -service student %s", file)
					doUpdateIdl(updateKitex)
					delete(changes, file)
				}
			}
		}
	}()

	// 保持程序运行
	done := make(chan bool)
	<-done
}

func doUpdateIdl(command string) {
	fmt.Println("kitex_server do ", command)

	// 将环境切换到当前目录
	changeToExecDir()

	// 创建一个 ls 命令
	cmd := exec.Command("/bin/sh", "-c", command)

	// 获取命令的输出和错误输出
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// 运行命令
	err := cmd.Run()
	if err != nil {
		log.Fatalf("命令执行失败: %s, %s", err.Error(), stderr.String())
	}
	// 打印输出
	//log.Println(out.String())
}

func changeToExecDir() {

	// 将当前工作目录更改为执行程序所在目录
	err := os.Chdir("/Users/xiaofeng/workplace_go/api-gateway/student_server/http_kitex")
	if err != nil {
		log.Fatalf("更改工作目录失败: %v", err)
	}

	// 获取更改后的当前工作目录
	_, err = os.Getwd()
	if err != nil {
		log.Fatalf("获取更改后的工作目录失败: %v", err)
	}
	//fmt.Printf("更改后的当前工作目录: %s\n", newCurrentDir)
}

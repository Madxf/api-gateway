## 开始：
```azure
1. git clone https://github.com/Madxf/api-gateway

## 开启客户端
go run ./student_client/http_hz_server

## 开启服务端
go run ./student_server/http_kitex

```

## feature

#### 20240506: 
1. 实现http泛化调用
#### 20240507: 
1. 实现json泛化调用


#### 20240526:
idl 热更新方案：  
1. 设置idl文件监控机制，当idl文件发生变化5s后，执行命令进行更新。  
2. 在 client 端，执行 hz update 命令进行更新  
3. 在 server 端，执行 kitex -module 命令进行更新  
4. 新增方法后，还需对业务逻辑进行处理。如果没有进行业务逻辑对处理，则需采取对应的返回方案。

#### 20240527:  
1. 完成 server 端的 idl 热更新  
2. 完成 client 端的 idl 热更新

#### 20240605:  
1. 新增对异常业务的处理

#### 20240610:  
1. 新增 client本地缓存 以优化访问速率

## 问题：  
1. idl 热更新后需要重启 client 才能生效。这也符合在 idl 更新后，需要编写相应业务逻辑的需求。

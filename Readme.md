# Go gRPC Microservices Project

This project consists of 3 microservices written in Go, communicating via gRPC. It also uses **gRPC-Gateway** to expose RESTful HTTP endpoints.

---

## 🧩 Features

- 3 Go-based services communicating via gRPC
- REST access via gRPC-Gateway
- Modular folder structure

---

## 🗂 Project Structure
```
gRPC Gateway
│   go.mod
│   go.sum
│   Makefile
│   Readme.md
│   script.sh
│   
├───api
│   ├───google
│   │   └───api
│   │           annotations.proto
│   │           http.proto
│   │           
│   ├───service1
│   │       service1.pb.go
│   │       service1.pb.gw.go
│   │       service1.proto
│   │       service1_grpc.pb.go
│   │       
│   ├───service2
│   │       service2.pb.go
│   │       service2.pb.gw.go
│   │       service2.proto
│   │       service2_grpc.pb.go
│   │       
│   └───service3
│           service3.pb.go
│           service3.pb.gw.go
│           service3.proto
│           service3_grpc.pb.go
│           
├───cmd
│   ├───client
│   │       main.go
│   │       
│   └───server
│           main.go
│           
└───internal
        service1.go
        service2.go
        service3.go
        
```

# Go gRPC Microservices Project

This project consists of 3 microservices written in Go, communicating via gRPC. It also uses **gRPC-Gateway** to expose RESTful HTTP endpoints.

---

## 🧩 Features

- 3 Go-based services communicating via gRPC
- REST access via gRPC-Gateway
- Modular folder structure

---

## 🗂 Project Structure

gRPC Gateway
│   go.mod
│   go.sum
│   Makefile
│   Readme.md
│   script.sh
│   structure.txt
│   
├───api
│   ├───google
│   │   └───api
│   │           annotations.proto
│   │           http.proto
│   │           
│   ├───monitoring_mgmt
│   │       monitoring_mgmt.pb.go
│   │       monitoring_mgmt.pb.gw.go
│   │       monitoring_mgmt.proto
│   │       monitoring_mgmt_grpc.pb.go
│   │       
│   ├───power_mgmt
│   │       power_mgmt.pb.go
│   │       power_mgmt.pb.gw.go
│   │       power_mgmt.proto
│   │       power_mgmt_grpc.pb.go
│   │       
│   └───script_mgmt
│           script_mgmt.pb.go
│           script_mgmt.pb.gw.go
│           script_mgmt.proto
│           script_mgmt_grpc.pb.go
│           
├───cmd
│   ├───client
│   │       main.go
│   │       
│   └───server
│           main.go
│           
└───internal
        monitoring_mgmt.go
        power_mgmt.go
        script_mgmt.go
        


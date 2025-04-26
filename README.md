# grpc-chat

# 🗨️ gRPC Chat Application in Go

This is a simple **gRPC-based bidirectional streaming chat server** and client built with Go.  

---

## 📄 Requirements

- Go 1.20+ installed
- `protoc` (Protocol Buffers Compiler)
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins installed

Install the plugins if you don't have them:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Setup
1. Clone the repository:
```bash
git clone git@github.com:t2013anurag/grpc-chat.git
cd grpc-chat
```

2. Generate Go code from the ```.proto``` file:
```bash
protoc --go_out=. --go-grpc_out=. chat.proto
```

## Usage

1. Start the server
```bash
go run server/main.go 
```

2. Start the client
```bash
go run client/main.go 
```

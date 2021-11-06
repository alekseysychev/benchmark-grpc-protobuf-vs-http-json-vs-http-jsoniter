### gRPC+Protobuf or JSON+HTTP or JSOUITER+HTTP?

This repository contains 2 equal APIs: gRPC using Protobuf and JSON over HTTP. The goal is to run benchmarks for 2 approaches and compare them. APIs have 1 endpoint to create user, containing validation of request. Request, validation and response are the same in 2 packages, so we're benchmarking only mechanism itself. Benchmarks also include response parsing.

### Requirements

- Go 1.16

### Run tests

Run benchmarks:

```bash
GO111MODULE=on go test -bench=. -benchmem
```

### Results

```bash
goos: linux
goarch: amd64
pkg: github.com/alekseysychev/benchmark-grpc-protobuf-vs-http-json-vs-http-jsoniter
cpu: Intel(R) Core(TM) i5-8300H CPU @ 2.30GHz
BenchmarkGRPCProtobuf-8            15624             79738 ns/op            7273 B/op        153 allocs/op
BenchmarkHTTPJSON-8                17318             73214 ns/op            9093 B/op        117 allocs/op
BenchmarkHTTPJSONIter-8            17762             65861 ns/op            9425 B/op        115 allocs/op
PASS
ok      github.com/alekseysychev/benchmark-grpc-protobuf-vs-http-json-vs-http-jsoniter  8.866s
```

They are almost the same, HTTP+JSON is a bit faster and has less allocs/op.

### CPU usage comparison

This will create an executable `benchmark-grpc-protobuf-vs-http-json.test` and the profile information will be stored in `grpcprotobuf.cpu` and `httpjson.cpu`:

```bash
GO111MODULE=on go test -bench=BenchmarkGRPCProtobuf -cpuprofile=_grpcprotobuf.cpu
GO111MODULE=on go test -bench=BenchmarkHTTPJSON -cpuprofile=_httpjson.cpu
GO111MODULE=on go test -bench=BenchmarkHTTPJSONIter -cpuprofile=_httpjsoniter.cpu
```

Check CPU usage per approach using:

```bash
go tool pprof _grpcprotobuf.cpu
go tool pprof _httpjson.cpu
go tool pprof _httpjsoniter.cpu
```

My results show that Protobuf consumes less ressources, around **30% less**.
But JsonIter less Protobuf, around **30% less**.

### gRPC definition

- Install [Go](https://golang.org/dl/)
- Install [Protocol Buffers](https://github.com/google/protobuf/releases)
- Install protoc plugin: `go get github.com/golang/protobuf/proto github.com/golang/protobuf/protoc-gen-go`

```bash
protoc --go_out=plugins=grpc:. grpc-protobuf/proto/api.proto
```

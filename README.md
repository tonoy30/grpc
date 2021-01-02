## Learning gRPC using golang

### Server

1. open a tcp connection
2. init gprc services struct
3. make a new grpc server
4. register service server
5. serve

```go
lis, err := net.Listen("tcp", ":9000")
s := chat.Server{}
grpcServer := grpc.NewServer()
chatpb.RegisterChatServiceServer(grpcServer, &s)
grpcServer.Serve(lis)
```

### Client

1. take grpc client connection
2. dial on opened tcp connection
3. register grpc client service
4. invoke rpc method

```go
var conn *grpc.ClientConn
conn, err := grpc.Dial(":9000", grpc.WithInsecure())
c := chatpb.NewChatServiceClient(conn)
greet, err := gs.Greeting(context.Background(), &greetpb.Greet{
Message: "Hello ",
})
```

### Protobuf

```protobuf
syntax = "proto3";

package dummy;

option go_package = "dummy/dummypb";

message Dummy {
  string dummyId = 1;  
}

service DummyService {
  // unary
  rpc InvokeDummy(Dummy) returns (Dummy) {}
  // server streaming
  rpc InvokeDummyStream(Dummy) returns (stream Dummy) {}
}
```

### Server Streaming API
> The client will send one message to the server and will receive many responses from the server,
> possibly an infinite number
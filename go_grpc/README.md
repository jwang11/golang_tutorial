# demo grpc usage


## create proto file hello_grpc.proto in pb/ folder
```
syntax = "proto3";
​
package service;
​
option go_package = ".;hello_grpc";
​
service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}
​
message HelloRequest {
  string name = 1;
}
​
message HelloReply {
  string message = 1;
}
```

## use compiler to generate service go file in service folder
```
protoc -I pb/ pb/hello_grpc.proto --go_out=plugins=grpc:service
```

##creaet server and client to use grpc service


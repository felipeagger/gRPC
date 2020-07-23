# gRPC
Demo gRPC Project with Golang and Python

Server and Client communication.

# ProtoBuf model

```
syntax = "proto3";
package user;

service UserService {
  rpc GetUser(UserRequest) returns (UserResponse) {}
}

message UserRequest {
  string username = 1;
}

message UserResponse {
  int64 id = 1;
  string name = 2;
  string username = 3;
  string avatarurl = 4;
  string location = 5;
  int64 repos = 6; 
}
```

# Running 

Golang running on Port 8200.

Python running on Port 8300.

```
//Server
make run

//client
make run_client
```

# Installation required:

ProtoBuff/Protoc
http://google.github.io/proto-lens/installing-protoc.html

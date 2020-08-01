# gRPC
Demo gRPC Project with Golang and Python

Server and Client communication.

**Golang com gRPC - Comunicação entre micro-serviços**

[![YouTube Video Explanation](http://img.youtube.com/vi/2b_4QA6D1qw/0.jpg)](http://www.youtube.com/watch?v=2b_4QA6D1qw "Golang com gRPC - Comunicação entre micro-serviços")

# Pontos positivos

- Estrutura RPC leve e de alto desempenho. Alem de usar HTTP/2 que é muito mais rápido que o HTTP/1.1 usado no REST padrão.

- Usa o protobuf para serializar dados, trafegando-os em binários.

- gRPC possui streaming de dados, e é bidirecional;

# Pontos Negativos

- Nao suporta comunicacao com Browsers (é mais indicado o uso entre servicos)

- Por trafegar os dados em binários, nao são legiveis aos humanos


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
  Statistics statistics = 6;
  repeated string listURLs = 7;
}

message Statistics {
  int64 followers = 1; 
  int64 following = 2; 
  int64 repos = 3; 
  int64 gists = 4; 
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

# Output

Golang 

```
2020/07/23 10:04:28 Response from server: 
id: 43504172 
name: "Filipe Alves" 
username: "felipeagger" 
avatarurl: "https://avatars0.githubusercontent.com/u/43504172?v=4" 
location: "Blumenau - SC / Brazil" 
statistics: {
  followers:8  
  following:4  
  repos:19  
  gists:1
}
listURLs: ["https://api.github.com/users/felipeagger", 
           "https://api.github.com/users/felipeagger/starred", 
           "https://api.github.com/users/felipeagger/repos"]
```

Python
```
<class 'user_pb2.UserResponse'>
id: 43504172
name: "Filipe Alves"
username: "felipeagger"
avatarurl: "https://avatars0.githubusercontent.com/u/43504172?v=4"
location: "Blumenau - SC / Brazil"
repos: 17
```

# Installation required:

ProtoBuff/Protoc
http://google.github.io/proto-lens/installing-protoc.html

# Links

https://developers.google.com/protocol-buffers

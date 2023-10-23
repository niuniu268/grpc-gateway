# gRPC-gateway example

![introduction](https://github.com/niuniu268/grpc-gateway/blob/master/img/Screenshot%202023-10-23%20at%2007.55.00.png?raw=true)

- Check protoc protoc-gen-go protoc-gen-go-grpc
- Create protoc file
- download: google/api/annotation and google/api/http
- create pb.go file
```
    protoc -I=proto \\n 
      --go_out=proto --go_opt=paths=source_relative \\n   
      --go-grpc_out=proto --go-grpc_opt=paths=source_relative \\n   
      helloworld/hello_world.proto
    
```

- get grpc and grpc/runtime (attention: v2 )

```
    go get -u google.golang.org/grpc
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2
    
    go get github.com/grpc-ecosystem/grpc-gateway/v2/runtime


```

- check the example

```

    curl -X POST -k http://localhost:8090/v1/example/echo -d '{"name": " hello"}'

```

package main

import Server "Gateway/Server"

//type server struct {
//	helloworldpb.UnimplementedGreeterServer
//}
//
//func NewServer() *server {
//	return &server{}
//}
//
//func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
//	return &helloworldpb.HelloReply{Message: in.Name + " world"}, nil
//}

func main() {

	Server.BuildupServer()

}

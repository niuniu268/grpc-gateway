package Test

import (
	helloworldpb "Gateway/proto/helloworld"
	"context"
	"testing"

	S "Gateway/Server"
)

func TestSayHello(t *testing.T) {

	server := S.NewServer()

	request := &helloworldpb.HelloRequest{
		Name: "niuniu",
	}

	response, err := server.SayHello(context.Background(), request)

	if err != nil {
		t.Errorf("SayHello failed: %v", err)
	}

	// Check if the response is as expected
	if response.Message != "niuniu world" {
		t.Errorf("Unexpected response: %v", response.Message)
	}

}

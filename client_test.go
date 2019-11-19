package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"go-micro-greeter/greeter"
	"testing"
)

func TestClient(t *testing.T) {
	// Create a new service
	service := micro.NewService(micro.Name("greeter.client"))
	// Initialise the client and parse command line flags
	service.Init()

	// Create new greeter client
	greeterClient := greeter.NewGreeterService("greeter", service.Client())

	// Call the greeter
	rsp, err := greeterClient.Hello(context.TODO(), &greeter.Request{Name: "John"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Msg)
}

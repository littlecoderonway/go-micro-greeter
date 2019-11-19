package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"go-micro-greeter/greeter"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *greeter.Request, rsp *greeter.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	greeter.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}

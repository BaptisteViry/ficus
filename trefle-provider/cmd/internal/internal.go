package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"ficus/branches/nursery/provider"

	"google.golang.org/grpc"

	"trefle-provider/handler/plantmanager"
	"trefle-provider/pkg/roundtripper"
)

const port = 8080

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var (
		pmh = plantmanager.NewHandler(
			&http.Client{
				Transport: &roundtripper.AuthRoundTripper{
					Next: http.DefaultTransport,
				},
			},
		)

		grpcServer = grpc.NewServer()
	)

	provider.RegisterPlantManagerServer(grpcServer, pmh)

	grpcServer.Serve(lis)
}

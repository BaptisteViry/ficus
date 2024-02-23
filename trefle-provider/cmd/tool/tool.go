package main

import (
	"context"
	"log"

	"ficus/branches/nursery/provider"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	plantManagerClient := provider.NewPlantManagerClient(conn)

	req := provider.FetchPlantsRequest{
		Query: "Pachira",
		Page: &provider.Page{
			Size:   10,
			Number: 1,
		},
	}

	plants, err := plantManagerClient.FetchPlants(ctx, &req)

	if err != nil {
		log.Fatalf("fail to fetch plants: %v", err)
	}

	log.Println(plants)
}

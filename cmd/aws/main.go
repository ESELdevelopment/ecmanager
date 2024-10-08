package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ESELDevelopment/ecmanager/internal/aws"
	"log"
)

func main() {
	ctx := context.Background()

	ecsService := aws.GetEcsService(ctx)

	clusters, err := ecsService.ListClusters(ctx)
	if err != nil {
		log.Fatalf("failed to list ECS clusters, %v", err)
	}

	for _, cluster := range clusters {
		detailsPointer, err := ecsService.DescribeClusters(ctx, cluster)
		if err != nil {
			log.Fatalf("failed retrieve details for cluster, %v with error %v", cluster, err)
		}
		detailsJson, _ := json.Marshal(*detailsPointer)
		fmt.Printf("Cluster: %v Details %v\n", cluster, string(detailsJson))
	}
}

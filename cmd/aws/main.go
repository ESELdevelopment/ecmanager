package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"ecmanager/internal/aws"
)

func main() {
	ctx := context.Background()

	awsClient, err := aws.GetAwsClient(ctx)

	if err != nil {
		log.Fatalf("failed to retrieve ECS client, %v", err)
	}
	clusters, err := awsClient.ListECSClusters(ctx)
	if err != nil {
		log.Fatalf("failed to list ECS clusters, %v", err)
	}

	for _, cluster := range clusters {
		detailsPointer, err := awsClient.DescribeECSCluster(ctx, cluster)
		if err != nil {
			log.Fatalf("failed retrieve details for cluster, %v with error %v", cluster, err)
		}
		detailsJson, _ := json.Marshal(*detailsPointer)
		fmt.Printf("Cluster: %v Details %v\n", cluster, string(detailsJson))
	}
}

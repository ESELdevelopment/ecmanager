package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

func (c *AwsClient) DescribeECSCluster(ctx context.Context, clusterName string) (*types.Cluster, error) {
	input := &ecs.DescribeClustersInput{
		Clusters: []string{clusterName},
	}
	result, err := c.ecsClient.DescribeClusters(ctx, input)

	return &result.Clusters[0], err
}

func (c *AwsClient) ListECSClusters(ctx context.Context) ([]string, error) {
	var clusters []string
	input := &ecs.ListClustersInput{}
	for {
		output, err := c.ecsClient.ListClusters(ctx, input)
		if err != nil {
			return nil, err
		}

		clusters = append(clusters, output.ClusterArns...)

		if output.NextToken == nil {
			break
		}
		input.NextToken = output.NextToken
	}
	return clusters, nil
}

func (c *AwsClient) TestMe() bool {
	return c.ecsClient != nil
}

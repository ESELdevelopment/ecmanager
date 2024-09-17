package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"net/url"
	"os"
	"sync"
)

var (
	once       = &sync.Once{}
	ecsService *ECSServiceImpl
)

type resolverV2 struct {
}

func (*resolverV2) ResolveEndpoint(ctx context.Context, params ecs.EndpointParameters) (
	smithyendpoints.Endpoint, error) {
	u := os.Getenv("AWS_ENDPOINT_URL")
	if u != "" {
		endpointUrl, err := url.Parse(u)
		if err != nil {
			return smithyendpoints.Endpoint{}, err
		}
		return smithyendpoints.Endpoint{
			URI: *endpointUrl,
		}, nil
	}
	return ecs.NewDefaultEndpointResolverV2().ResolveEndpoint(ctx, params)
}

func GetEcsService(ctx context.Context) ECSService {
	once.Do(func() {
		cfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			return
		}
		ecsClient := ecs.NewFromConfig(cfg, func(o *ecs.Options) {
			o.EndpointResolverV2 = &resolverV2{}
		})

		ecsService = &ECSServiceImpl{client: ecsClient}
	})
	return ecsService
}

type ECSService interface {
	DescribeClusters(ctx context.Context, clusterName string) (*ecs.DescribeClustersOutput, error)
	ListClusters(ctx context.Context) ([]string, error)
}

type ECSServiceImpl struct {
	client *ecs.Client
}

func (e *ECSServiceImpl) DescribeClusters(ctx context.Context, clusterName string) (*ecs.DescribeClustersOutput, error) {
	input := &ecs.DescribeClustersInput{
		Clusters: []string{clusterName},
	}

	return e.client.DescribeClusters(ctx, input)
}

func (c *ECSServiceImpl) ListClusters(ctx context.Context) ([]string, error) {
	var clusters []string
	input := &ecs.ListClustersInput{}
	for {
		output, err := c.client.ListClusters(ctx, input)
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

package aws

import (
	"context"
	"net/url"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
)

var lock = &sync.Mutex{}

type AwsClient struct {
	ecsClient *ecs.Client
}

var client *AwsClient

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

func GetAwsClient(ctx context.Context) (*AwsClient, error) {
	if client == nil {
		lock.Lock()
		defer lock.Unlock()
		// avoid creating multiple clients on multiple goroutines
		if client == nil {
			cfg, err := config.LoadDefaultConfig(context.TODO())
			if err != nil {
				return nil, err
			}
			client = &AwsClient{
				ecsClient: ecs.NewFromConfig(cfg, func(o *ecs.Options) {
					o.EndpointResolverV2 = &resolverV2{}
				}),
			}
		}
	}
	return client, nil
}

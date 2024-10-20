package sts

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"net/url"
	"os"
	"sync"
)

type resolverV2 struct {
}

var (
	once       = &sync.Once{}
	stsService *ServiceImpl
)

func (r *resolverV2) ResolveEndpoint(ctx context.Context, params sts.EndpointParameters) (
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
	return sts.NewDefaultEndpointResolverV2().ResolveEndpoint(ctx, params)
}

func GetService(ctx context.Context) Service {
	once.Do(func() {
		cfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			return
		}
		stsClient := sts.NewFromConfig(cfg, func(o *sts.Options) {
			o.EndpointResolverV2 = &resolverV2{}
		})
		stsService = &ServiceImpl{client: stsClient}
	})
	return stsService
}

type Service interface {
	GetCallerIdentity(ctx context.Context) (*sts.GetCallerIdentityOutput, error)
}

type ServiceImpl struct {
	client *sts.Client
}

func (s *ServiceImpl) GetCallerIdentity(ctx context.Context) (*sts.GetCallerIdentityOutput, error) {
	return s.client.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
}

package ecs_test

import (
	"context"
	"fmt"
	"github.com/ESELDevelopment/ecmanager/internal/aws/ecs"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	ecsSdk "github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	smithyMiddleware "github.com/aws/smithy-go/middleware"
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"testing"
	"unsafe"
)

func TestGetEcsService(t *testing.T) {
	ecsService := ecs.GetService(context.TODO())
	assert.NotNilf(t, ecsService, "GetEcsService() = %v, want not nil", ecsService)
	ecsService2 := ecs.GetService(context.TODO())
	assert.Equal(t, ecsService, ecsService2)
}

func TestECSServiceImpl_DescribeClusters(t *testing.T) {
	type args struct {
		ctx                context.Context
		withApiOptionsFunc func(stack *smithyMiddleware.Stack) error
	}

	cases := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "Retrieve cluster information",
			args: args{
				ctx: context.Background(),
				withApiOptionsFunc: func(stack *smithyMiddleware.Stack) error {
					return stack.Finalize.Add(
						smithyMiddleware.FinalizeMiddlewareFunc(
							"DescribeClustersMock",
							func(ctx context.Context, input smithyMiddleware.FinalizeInput, handler smithyMiddleware.FinalizeHandler) (smithyMiddleware.FinalizeOutput, smithyMiddleware.Metadata, error) {
								return smithyMiddleware.FinalizeOutput{
									Result: &ecsSdk.DescribeClustersOutput{
										Clusters:       []types.Cluster{{ClusterName: aws.String(`clusterName`)}},
										Failures:       nil,
										ResultMetadata: smithyMiddleware.Metadata{},
									},
								}, smithyMiddleware.Metadata{}, nil
							},
						), smithyMiddleware.Before)
				}},
			want:    []types.Cluster{{ClusterName: aws.String(`clusterName`)}},
			wantErr: false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			os.Unsetenv("AWS_ACCESS_KEY_ID")
			os.Unsetenv("AWS_SECRET_ACCESS_KEY")
			cfg, err := config.LoadDefaultConfig(
				tt.args.ctx,
				config.WithRegion("eu-central-1"),
				config.WithAPIOptions([]func(*smithyMiddleware.Stack) error{tt.args.withApiOptionsFunc}))
			if err != nil {
				t.Fatal(err)
			}

			client := ecsSdk.NewFromConfig(cfg)
			ecsService := ecs.ServiceImpl{}
			setFieldValue(&ecsService, "client", client)

			response, e := ecsService.DescribeClusters(tt.args.ctx, "clusterName")

			if tt.wantErr {
				assert.NotNilf(t, e, "DescribeClusters() error = %v, wantErr %v", e, tt.wantErr)
			} else {
				assert.Equal(t, tt.want, response.Clusters)
			}
		})
	}
}

func TestECSServiceImpl_ListClusters(t *testing.T) {
	type args struct {
		ctx                context.Context
		withApiOptionsFunc func(stack *smithyMiddleware.Stack) error
	}

	cases := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "List clusters with pagination",
			args: args{
				ctx: context.Background(),
				withApiOptionsFunc: func(stack *smithyMiddleware.Stack) error {
					var customInitialize = smithyMiddleware.InitializeMiddlewareFunc("customInitialize", func(
						ctx context.Context, in smithyMiddleware.InitializeInput, next smithyMiddleware.InitializeHandler,
					) (
						out smithyMiddleware.InitializeOutput, metadata smithyMiddleware.Metadata, err error,
					) {
						ctx = SetCustomKey(ctx, in.Parameters)

						return next.HandleInitialize(ctx, in)
					})

					_ = stack.Initialize.Add(customInitialize, smithyMiddleware.Before)

					return stack.Finalize.Add(
						smithyMiddleware.FinalizeMiddlewareFunc(
							"ListClustersMock",
							func(ctx context.Context, input smithyMiddleware.FinalizeInput, handler smithyMiddleware.FinalizeHandler) (smithyMiddleware.FinalizeOutput, smithyMiddleware.Metadata, error) {
								request := GetCustomKey(ctx).(*ecsSdk.ListClustersInput)
								if request.NextToken != nil {
									return smithyMiddleware.FinalizeOutput{
										Result: &ecsSdk.ListClustersOutput{
											ClusterArns: []string{"clusterArn2"},
											NextToken:   nil,
										},
									}, smithyMiddleware.Metadata{}, nil
								}

								return smithyMiddleware.FinalizeOutput{
									Result: &ecsSdk.ListClustersOutput{
										ClusterArns: []string{"clusterArn1"},
										NextToken:   aws.String("nextToken"),
									},
								}, smithyMiddleware.Metadata{}, nil
							},
						), smithyMiddleware.Before)
				}},
			want:    []string{"clusterArn1", "clusterArn2"},
			wantErr: false,
		},
		{
			name: "Fails on Error",
			args: args{
				ctx: context.Background(),
				withApiOptionsFunc: func(stack *smithyMiddleware.Stack) error {
					return stack.Finalize.Add(
						smithyMiddleware.FinalizeMiddlewareFunc(
							"ListClustersMock",
							func(ctx context.Context, input smithyMiddleware.FinalizeInput, handler smithyMiddleware.FinalizeHandler) (smithyMiddleware.FinalizeOutput, smithyMiddleware.Metadata, error) {
								return smithyMiddleware.FinalizeOutput{}, smithyMiddleware.Metadata{}, fmt.Errorf("error")
							},
						), smithyMiddleware.Before)
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			os.Unsetenv("AWS_ACCESS_KEY_ID")
			os.Unsetenv("AWS_SECRET_ACCESS_KEY")
			cfg, err := config.LoadDefaultConfig(
				tt.args.ctx,
				config.WithRegion("eu-central-1"),
				config.WithAPIOptions([]func(*smithyMiddleware.Stack) error{tt.args.withApiOptionsFunc}))
			if err != nil {
				t.Fatal(err)
			}

			client := ecsSdk.NewFromConfig(cfg)
			ecsService := ecs.ServiceImpl{}
			setFieldValue(&ecsService, "client", client)

			list, e := ecsService.ListClusters(tt.args.ctx)

			if tt.wantErr {
				assert.NotNilf(t, e, "ListClusters() error = %v, wantErr %v", e, tt.wantErr)
				assert.Errorf(t, e, "error")
			} else {
				assert.ElementsMatch(t, tt.want, list)
			}
		})
	}
}

func setFieldValue(target any, fieldName string, value any) {
	rv := reflect.ValueOf(target)
	for rv.Kind() == reflect.Ptr && !rv.IsNil() {
		rv = rv.Elem()
	}
	if !rv.CanAddr() {
		panic("target must be addressable")
	}
	if rv.Kind() != reflect.Struct {
		panic(fmt.Sprintf(
			"unable to set the '%s' field value of the type %T, target must be a struct",
			fieldName,
			target,
		))
	}
	rf := rv.FieldByName(fieldName)

	reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem().Set(reflect.ValueOf(value))
}

type customKey struct{}

func GetCustomKey(ctx context.Context) (v interface{}) {
	v = smithyMiddleware.GetStackValue(ctx, customKey{})
	return v
}

func SetCustomKey(ctx context.Context, value interface{}) context.Context {
	return smithyMiddleware.WithStackValue(ctx, customKey{}, value)
}

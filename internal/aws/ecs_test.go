package aws_test

import (
	"context"
	"ecmanager/internal/aws"
	"fmt"
	awsSdk "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	"os"
	"reflect"
	"testing"
	"unsafe"
)

func TestECSService_DescribeClusters(t *testing.T) {
	type args struct {
		ctx                context.Context
		withApiOptionsFunc func(stack *middleware.Stack) error
	}

	cases := []struct {
		name    string
		args    args
		want    error
		wantErr bool
	}{
		{
			name: "Retrieve cluster information",
			args: args{
				ctx: context.Background(),
				withApiOptionsFunc: func(stack *middleware.Stack) error {
					return stack.Finalize.Add(
						middleware.FinalizeMiddlewareFunc(
							"DescribeClustersMock",
							func(ctx context.Context, input middleware.FinalizeInput, handler middleware.FinalizeHandler) (middleware.FinalizeOutput, middleware.Metadata, error) {
								return middleware.FinalizeOutput{

									Result: &ecs.DescribeClustersOutput{
										Clusters:       []types.Cluster{{ClusterName: awsSdk.String(`clusterName`)}},
										Failures:       nil,
										ResultMetadata: middleware.Metadata{},
									},
								}, middleware.Metadata{}, nil
							},
						), middleware.Before)
				}},
			want:    nil,
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
				config.WithAPIOptions([]func(*middleware.Stack) error{tt.args.withApiOptionsFunc}))
			if err != nil {
				t.Fatal(err)
			}

			client := ecs.NewFromConfig(cfg)
			ecsService := aws.ECSServiceImpl{}
			setFieldValue(&ecsService, "client", client)

			_, err = ecsService.DescribeClusters("clusterName")
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

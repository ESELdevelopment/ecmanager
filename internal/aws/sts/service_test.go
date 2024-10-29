package sts_test

import (
	"context"
	"fmt"
	"github.com/ESELDevelopment/ecmanager/internal/aws/sts"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	stsSdk "github.com/aws/aws-sdk-go-v2/service/sts"
	smithyMiddleware "github.com/aws/smithy-go/middleware"
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"testing"
	"unsafe"
)

func TestGetService(t *testing.T) {
	service := sts.GetService(context.TODO())
	assert.NotNil(t, service)
	service2 := sts.GetService(context.TODO())
	assert.Equal(t, service, service2)
}

func TestServiceImpl_GetCallerIdentity(t *testing.T) {
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
			name: "Retrieve caller identity",
			args: args{
				ctx: context.Background(),
				withApiOptionsFunc: func(stack *smithyMiddleware.Stack) error {
					return stack.Finalize.Add(
						smithyMiddleware.FinalizeMiddlewareFunc(
							"GetCallerIdentityMock",
							func(ctx context.Context, input smithyMiddleware.FinalizeInput, handler smithyMiddleware.FinalizeHandler) (smithyMiddleware.FinalizeOutput, smithyMiddleware.Metadata, error) {
								return smithyMiddleware.FinalizeOutput{
									Result: &stsSdk.GetCallerIdentityOutput{
										Account: aws.String(`account`),
										Arn:     aws.String(`arn`),
										UserId:  aws.String(`userId`),
									},
								}, smithyMiddleware.Metadata{}, nil
							},
						), smithyMiddleware.Before)
				}},
			want: &stsSdk.GetCallerIdentityOutput{
				Account: aws.String(`account`),
				Arn:     aws.String(`arn`),
				UserId:  aws.String(`userId`),
			},
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

			client := stsSdk.NewFromConfig(cfg)
			serviceImpl := sts.ServiceImpl{}
			setFieldValue(&serviceImpl, "client", client)

			response, e := serviceImpl.GetCallerIdentity(tt.args.ctx)

			if tt.wantErr {
				assert.NotNilf(t, e, "DescribeClusters() error = %v, wantErr %v", e, tt.wantErr)
			} else {
				assert.Equal(t, tt.want, response)
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

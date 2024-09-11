package aws_test

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"ecmanager/internal/aws"

	"github.com/stretchr/testify/assert"
)

type MockEcsClient struct{}

func (m *MockEcsClient) ListClustersInput() string {
	return "Hello"
}

func TestDescribeClustersInputTest(t *testing.T) {
	client := &aws.AwsClient{}
	setFieldValue(client, "ecsClient", MockEcsClient{})
	result := client.TestMe()
	assert.True(t, result)
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

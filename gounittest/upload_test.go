package main

import (
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockS3 struct {
	mock.Mock
}

func (m *MockS3) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*s3.PutObjectOutput), args.Error(1)
}

func TestUpload_Mock(t *testing.T) {
	m := new(MockS3)
	u := Uploader{
		svc: m,
	}
	// m.On("PutObject", mock.Anything).Return(&s3.PutObjectOutput{}, nil)
	m.On("PutObject", &s3.PutObjectInput{
		Body:    aws.ReadSeekCloser(strings.NewReader("c:\\HappyFace.jpg")),
		Bucket:  aws.String("examplebucket"),
		Key:     aws.String("HappyFace.jpg"),
		Tagging: aws.String("key1=value1&key2=value2"),
	}).Return(&s3.PutObjectOutput{}, nil)
	err := u.Upload()
	assert.NoError(t, err) // expect no error on third party

	m.AssertExpectations(t)
	// test PutObject func to expect PutObjectInput to receive specific object args and return specific value
}

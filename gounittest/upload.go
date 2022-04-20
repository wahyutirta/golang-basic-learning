package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AwsUploader interface {
	PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error)
}

type Uploader struct {
	svc AwsUploader
}

func (u *Uploader) Upload() error {
	input := &s3.PutObjectInput{
		Body:    aws.ReadSeekCloser(strings.NewReader("c:\\HappyFace.jpg")),
		Bucket:  aws.String("examplebucket"),
		Key:     aws.String("HappyFace.jpg"),
		Tagging: aws.String("key1=value1&key2=value2"),
	}
	result, err := u.svc.PutObject(input)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

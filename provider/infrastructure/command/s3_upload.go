package command

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

type BucketS3UploadObject struct {
	object *s3.S3
}

func NewBucketS3UploadObject(object *s3.S3) *BucketS3UploadObject {
	return &BucketS3UploadObject{object: object}
}

// Use of the command
func (b *BucketS3UploadObject) Use() string {
	return "s3:put"
}

// Example of the command
func (b *BucketS3UploadObject) Example() string {
	return "s3:put"
}

// Short description about the command
func (b *BucketS3UploadObject) Short() string {
	return "Upload items in S3 Bucket"
}

// Run the command with the args given by the caller
func (b *BucketS3UploadObject) Run(args []string) {
	bucket := os.Getenv("BUCKET_NAME")

	file, err := os.Open(args[0])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	input := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		ACL:    aws.String("private"),
		Key:    aws.String(fmt.Sprintf("temp/%s", args[0])),
		Body:   file,
		//ContentType:        aws.String(http.DetectContentType(message)),
		ContentDisposition: aws.String("attachment"),
	}

	aa, err := b.object.PutObject(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(aa.GoString())
	fmt.Println("Success Upload Object")
}

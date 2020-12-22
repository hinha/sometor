package command

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

type BucketS3List struct {
	object *s3.S3
}

func NewBucketS3List(object *s3.S3) *BucketS3List {
	return &BucketS3List{object: object}
}

// Use of the command
func (b *BucketS3List) Use() string {
	return "s3:list"
}

// Example of the command
func (b *BucketS3List) Example() string {
	return "s3:list"
}

// Short description about the command
func (b *BucketS3List) Short() string {
	return "Lists the items in the specified S3 Bucket"
}

// Run the command with the args given by the caller
func (b *BucketS3List) Run(args []string) {
	// Get the list of items
	bucket := os.Getenv("BUCKET_NAME")
	resp, err := b.object.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
	if err != nil {
		panic(err)
	}
	fmt.Println("Found", len(resp.Contents), "items in bucket", bucket)
	fmt.Println("")
}

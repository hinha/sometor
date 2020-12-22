package command

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

type BucketS3ListObject struct {
	object *s3.S3
}

func NewBucketS3ListObject(object *s3.S3) *BucketS3ListObject {
	return &BucketS3ListObject{object: object}
}

// Use of the command
func (b *BucketS3ListObject) Use() string {
	return "s3:list_object"
}

// Example of the command
func (b *BucketS3ListObject) Example() string {
	return "s3:list_object"
}

// Short description about the command
func (b *BucketS3ListObject) Short() string {
	return "Lists all object the items in S3 Bucket"
}

// Run the command with the args given by the caller
func (b *BucketS3ListObject) Run(args []string) {
	// Get the list of items
	bucket := os.Getenv("BUCKET_NAME")

	resp, err := b.object.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
	if err != nil {
		panic(err)
	}

	for _, item := range resp.Contents {
		fmt.Println("Name:         ", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:         ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
	}

	fmt.Println("Found", len(resp.Contents), "items in bucket", bucket)
	fmt.Println("")
}

package command

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

type BucketS3DownloadObject struct {
	object  *s3.S3
	session *session.Session
}

func NewBucketS3DownloadObject(object *s3.S3, sess *session.Session) *BucketS3DownloadObject {
	return &BucketS3DownloadObject{object: object, session: sess}
}

// Use of the command
func (b *BucketS3DownloadObject) Use() string {
	return "s3:download"
}

// Example of the command
func (b *BucketS3DownloadObject) Example() string {
	return "s3:download"
}

// Short description about the command
func (b *BucketS3DownloadObject) Short() string {
	return "Download items in S3 Bucket"
}

func (b *BucketS3DownloadObject) Run(args []string) {

	basePath, _ := os.Getwd()
	path := fmt.Sprintf("%s/temp", basePath)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	file, err := os.Create(args[0])
	if err != nil {
		panic(fmt.Sprintf("Unable to open file %q, %v", args[0], err))
	}

	defer file.Close()

	downloader := s3manager.NewDownloader(b.session)
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(os.Getenv("BUCKET_NAME")),
			Key:    aws.String(args[0]),
		})

	if err != nil {
		panic(fmt.Sprintf("Unable to download item %q, %v", args[0], err))
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}

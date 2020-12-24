package adapter

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

type S3 struct {
	object     *s3.S3
	session    *session.Session
	bucketName string
}

func AdaptS3(object *s3.S3, awSession *session.Session) *S3 {
	bucket := os.Getenv("BUCKET_NAME")
	return &S3{object: object, session: awSession, bucketName: bucket}
}

func (s *S3) PutObject(pathObject string, file *os.File) error {
	defer file.Close()

	input := &s3.PutObjectInput{
		Bucket:             aws.String(s.bucketName),
		ACL:                aws.String("private"),
		Key:                aws.String(pathObject),
		Body:               file,
		ContentDisposition: aws.String("attachment"),
	}

	aa, err := s.object.PutObject(input)
	if err != nil {
		return err
	}

	fmt.Println(aa.GoString())
	fmt.Println("Success Upload Object")

	return nil
}

// DownloadObject path by default is temp/filename
func (s *S3) DownloadObject(pathObject string) (string, error) {
	basePath, _ := os.Getwd()
	path := fmt.Sprintf("%s/temp", basePath)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	file, err := os.Create(pathObject)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Unable to open file %q, %v", pathObject, err))
	}

	defer file.Close()

	downloader := s3manager.NewDownloader(s.session)
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(s.bucketName),
			Key:    aws.String(pathObject),
		})

	if err != nil {
		return "", errors.New(fmt.Sprintf("Unable to download item %q, %v", pathObject, err))
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
	return file.Name(), nil
}

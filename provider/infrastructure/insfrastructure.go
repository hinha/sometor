package infrastructure

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hinha/sometor/provider"
	"github.com/hinha/sometor/provider/infrastructure/command"
	"os"
)

type Infrastructure struct {
	awsS3 struct {
		Object  *s3.S3
		Session *session.Session
	}
}

// Fabricate infrastructure interface for kalkula
func Fabricate() (*Infrastructure, error) {
	i := &Infrastructure{}

	creeds := credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")
	_, err := creeds.Get()
	if err != nil {
		return nil, err
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: creeds,
	})
	if err != nil {
		return nil, err
	}

	i.awsS3.Object = s3.New(sess)
	i.awsS3.Session = sess

	return i, nil
}

func (i *Infrastructure) FabricateCommand(cmd provider.Command) error {
	cmd.InjectCommand(
		command.NewBucketS3List(i.awsS3.Object),
		command.NewBucketS3ListObject(i.awsS3.Object),
		command.NewBucketS3UploadObject(i.awsS3.Object),
	)

	return nil
}

func (i *Infrastructure) Close() {

}

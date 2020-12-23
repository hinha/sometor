package infrastructure

import (
	"database/sql"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	_ "github.com/go-sql-driver/mysql" // Import mysql driver
	"github.com/gocelery/gocelery"
	"github.com/gomodule/redigo/redis"
	"github.com/hinha/sometor/provider"
	"github.com/hinha/sometor/provider/infrastructure/adapter"
	"github.com/hinha/sometor/provider/infrastructure/command"
	"os"
	"sync"
	"time"
)

type Infrastructure struct {
	mysqlMutex  *sync.Once
	mysqlDB     *sql.DB
	mysqlConfig struct {
		username string
		password string
		host     string
		port     string
		dbname   string
	}
	celeryClient struct {
		redis *redis.Pool
	}
	redisConfig struct {
		host string
	}

	awsS3Config struct {
		accessKeyId     string
		accessKeySecret string
		region          string
	}
}

// Fabricate infrastructure interface for sometor
func Fabricate() (*Infrastructure, error) {
	i := &Infrastructure{
		mysqlMutex: &sync.Once{},
	}

	i.mysqlConfig.host = os.Getenv("MYSQL_HOST")
	i.mysqlConfig.username = os.Getenv("MYSQL_USERNAME")
	i.mysqlConfig.password = os.Getenv("MYSQL_PASSWORD")
	i.mysqlConfig.dbname = os.Getenv("MYSQL_DATABASE")
	i.mysqlConfig.port = os.Getenv("MYSQL_PORT")

	i.redisConfig.host = os.Getenv("URI_REDIS_HOST")

	i.awsS3Config.accessKeyId = os.Getenv("AWS_ACCESS_KEY_ID")
	i.awsS3Config.accessKeySecret = os.Getenv("AWS_SECRET_ACCESS_KEY")
	i.awsS3Config.region = os.Getenv("AWS_REGION")

	return i, nil
}

func (i *Infrastructure) FabricateCommand(cmd provider.Command) error {
	s3Object, err := i.S3Object()
	if err != nil {
		return err
	}
	s3Session, err := i.AwSession()
	if err != nil {
		return err
	}

	cmd.InjectCommand(
		command.NewBucketS3List(s3Object),
		command.NewBucketS3ListObject(s3Object),
		command.NewBucketS3UploadObject(s3Object),
		command.NewBucketS3DownloadObject(s3Object, s3Session),
	)

	return nil
}

func (i *Infrastructure) Celery() provider.CeleryClient {
	redisPool := &redis.Pool{
		MaxIdle:     3,                 // maximum number of idle connections in the pool
		MaxActive:   0,                 // maximum number of connections allocated by the pool at a given time
		IdleTimeout: 240 * time.Second, // close connections after remaining idle for this duration
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(i.redisConfig.host)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	// initialize celery client
	cli, _ := gocelery.NewCeleryClient(
		gocelery.NewRedisBroker(redisPool),
		&gocelery.RedisCeleryBackend{Pool: redisPool},
		3,
	)

	return adapter.AdaptCelery(cli)
}

// MYSQL provide mysql interface
func (i *Infrastructure) MYSQL() (*sql.DB, error) {
	i.mysqlMutex.Do(func() {
		db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true",
			i.mysqlConfig.username,
			i.mysqlConfig.password,
			i.mysqlConfig.host,
			i.mysqlConfig.port,
			i.mysqlConfig.dbname))

		i.mysqlDB = db
	})

	return i.mysqlDB, nil
}

func (i *Infrastructure) Close() {
	if i.mysqlDB != nil {
		_ = i.mysqlDB.Close()
	}
}

func (i *Infrastructure) DB() (provider.DB, error) {
	db, err := i.MYSQL()
	if err != nil {
		return nil, err
	}

	return adapter.AdaptSQL(db), nil
}

func (i *Infrastructure) S3Object() (*s3.S3, error) {
	awSession, err := i.AwSession()
	if err != nil {
		return nil, err
	}
	return s3.New(awSession), nil
}

func (i *Infrastructure) AwSession() (*session.Session, error) {
	creeds := credentials.NewStaticCredentials(i.awsS3Config.accessKeyId, i.awsS3Config.accessKeySecret, "")
	_, err := creeds.Get()
	if err != nil {
		return nil, err
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(i.awsS3Config.region),
		Credentials: creeds,
	})
	if err != nil {
		return nil, err
	}

	return sess, nil
}

func (i *Infrastructure) S3() (provider.S3Management, error) {
	sess, err := i.AwSession()
	if err != nil {
		return nil, err
	}

	s3Ob, err := i.S3Object()
	if err != nil {
		return nil, err
	}

	return adapter.AdaptS3(s3Ob, sess), nil
}

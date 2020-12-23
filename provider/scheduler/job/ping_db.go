package job

import (
	"fmt"
	"github.com/gocraft/work"
)

type PingDB struct {
}

func NewPingDB() *PingDB {
	return &PingDB{}
}

func (p *PingDB) JobName() string {
	return "ping_db"
}

func (p *PingDB) JobTime() string {
	return "@every 0h1m0s"
}

func (p *PingDB) JobMiddleware(job *work.Job, next work.NextMiddlewareFunc) error {
	fmt.Println("Starting job: ", job.Name)
	return next()
}

func (p *PingDB) Retry() uint {
	return 3
}

func (p *PingDB) JobFunc(context *work.Job) error {
	fmt.Println("Successfully pinging mysql.")
	return nil
}

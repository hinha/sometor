package adapter

import (
	"github.com/gocelery/gocelery"
	"time"
)

type Celery struct {
	client *gocelery.CeleryClient
}

func AdaptCelery(client *gocelery.CeleryClient) *Celery {
	return &Celery{client: client}
}

func (c *Celery) Task(taskName string, args interface{}) (*gocelery.AsyncResult, error) {
	// run task
	asyncResult, err := c.client.Delay(taskName, args)
	if err != nil {
		return nil, err
	}
	return asyncResult, nil
}

func (c *Celery) Result(result *gocelery.AsyncResult) (interface{}, error) {
	// get results from backend with timeout
	res, err := result.Get(10 * time.Second)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Celery) GetTaskResult(taskName string, timeout int, args ...interface{}) (interface{}, error) {
	asyncResult, err := c.client.Delay(taskName, args)
	if err != nil {
		return nil, err
	}

	res, err := asyncResult.Get(time.Duration(timeout) * time.Minute)
	if err != nil {
		return nil, err
	}

	return res, nil
}

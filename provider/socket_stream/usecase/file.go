package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"io/ioutil"
	"os"
	"time"
)

type FileTwitter struct {
}

func (t *FileTwitter) Perform(_ context.Context, media, fileUser string, lastMod time.Time, userProvider provider.StreamSequence) ([]byte, time.Time, *entity.ApplicationError) {
	//var twitterData entity.TwitterResult
	//var portData []byte

	path, _ := os.Getwd()
	filename := fmt.Sprintf("%s/temp/account-%s-%s.json", path, fileUser, media)
	p, mod, err := t.readFileIfModified(filename, lastMod)
	if err != nil {
		return nil, lastMod, err
	}

	return p, mod, nil
}

func (t *FileTwitter) readFileIfModified(filename string, lastMod time.Time) ([]byte, time.Time, *entity.ApplicationError) {
	fi, err := os.Stat(filename)
	if err != nil {
		return nil, lastMod, &entity.ApplicationError{
			Err: []error{errors.New("collecting data")},
		}
	}
	if !fi.ModTime().After(lastMod) {
		return nil, lastMod, nil
	}

	p, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fi.ModTime(), &entity.ApplicationError{
			Err: []error{err},
		}
	}
	return p, fi.ModTime(), nil
}

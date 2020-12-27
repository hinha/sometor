package usecase

import (
	"context"
	"fmt"
	"github.com/hinha/sometor/provider"
	"io/ioutil"
	"os"
	"time"
)

type FileTwitter struct {
}

func (t *FileTwitter) Perform(_ context.Context, media, fileUser string, lastMod time.Time, userProvider provider.StreamSequence) ([]byte, time.Time, error) {
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

func (t *FileTwitter) readFileIfModified(filename string, lastMod time.Time) ([]byte, time.Time, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		return nil, lastMod, err
	}
	if !fi.ModTime().After(lastMod) {
		return nil, lastMod, nil
	}

	p, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fi.ModTime(), err
	}
	return p, fi.ModTime(), nil
}

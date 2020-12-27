package socmed

import (
	"context"
	"errors"
	"fmt"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"github.com/hinha/sometor/provider/socmed/api"
	"github.com/hinha/sometor/provider/socmed/usecase"
	"net/http"
	"strings"
)

type KeywordStream struct {
	providerKeyword provider.StreamKeyword
}

func FabricateKeyword(providerKeyword provider.StreamKeyword) *KeywordStream {
	return &KeywordStream{providerKeyword: providerKeyword}
}

func (k *KeywordStream) FabricateAPI(engine provider.APIEngine) {
	engine.InjectAPI(api.NewListStreamKeyword(k))
	engine.InjectAPI(api.NewCreateStreamKeyword(k))
}

func (k *KeywordStream) StreamKeywordList(ctx context.Context, ID string) ([]entity.StreamSequenceInitTable, *entity.ApplicationError) {
	listKeyword := usecase.ListStream{}
	return listKeyword.Perform(ctx, ID, k.providerKeyword)
}

func (k *KeywordStream) StreamKeywordCreate(ctx context.Context, request entity.StreamSequenceInsertable) (entity.StreamSequenceInsertable, *entity.ApplicationError) {
	var mediaValid bool

	fmt.Println(request)

	switch request.Media {
	case "twitter", "instagram", "facebook":
		mediaValid = true
		break
	default:
		mediaValid = false
		break
	}

	if mediaValid == false {
		return entity.StreamSequenceInsertable{}, &entity.ApplicationError{
			Err:        []error{errors.New("media not found")},
			HTTPStatus: http.StatusNotFound,
		}
	}

	if request.Type == "account" {
		request.Keyword = strings.ReplaceAll(request.Keyword, "@", "")
	} else if request.Type == "hashtag" {
		request.Keyword = strings.ReplaceAll(request.Keyword, "#", "")
	}

	result, err := k.providerKeyword.CreateOrFindStreamKeyword(ctx, request)
	if err != nil {
		return result, err
	}

	return result, nil
}
package provider

import (
	"context"
	"github.com/hinha/sometor/entity"
)

type SocmedKeywordAPI interface {
	StreamKeywordList(ctx context.Context, ID string) ([]entity.StreamSequenceInitTable, *entity.ApplicationError)
	StreamKeywordCreate(ctx context.Context, request entity.StreamSequenceInsertable) (entity.StreamSequenceInsertable, *entity.ApplicationError)
}

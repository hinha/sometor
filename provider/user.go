package provider

import (
	"context"
	"github.com/hinha/sometor/entity"
)

type StreamSequence interface {
	FindByUserID(ctx context.Context, ID string) (entity.StreamSequenceInitTable, *entity.ApplicationError)
	FindAllUser(ctx context.Context) ([]entity.StreamSequenceInitTable, *entity.ApplicationError)
}

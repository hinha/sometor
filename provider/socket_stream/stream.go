package socket_stream

import (
	"context"
	"github.com/hinha/sometor/entity"
	"github.com/hinha/sometor/provider"
	"github.com/hinha/sometor/provider/socket_stream/socket"
	"github.com/hinha/sometor/provider/socket_stream/usecase"
	"time"
)

type SocketStream struct {
	userProvider provider.StreamSequence
}

func Fabricate(userProvider provider.StreamSequence) *SocketStream {
	return &SocketStream{userProvider: userProvider}
}

func (s *SocketStream) FabricateSocket(engine provider.SocketEngine) {
	engine.InjectSocket(socket.NewTwitterSocket(s))
	engine.InjectSocket(socket.NewTwitterDemoWeb(s))
	engine.InjectSocket(socket.NewInstagramSocket(s))
	engine.InjectSocket(socket.NewInstagramDemoWeb(s))
}

func (s *SocketStream) FileReader(ctx context.Context, lastMod time.Time, media string, FileUser string) ([]byte, time.Time, *entity.ApplicationError) {
	files := usecase.FileTwitter{}
	return files.Perform(ctx, media, FileUser, lastMod, s.userProvider)
}

func (s *SocketStream) Writer(ctx context.Context, ID string, FileUser string) {

}

func (s *SocketStream) UserValid(ctx context.Context, ID, keyword, media string) (entity.UserAccountSelectable, *entity.ApplicationError) {
	valid := usecase.UserValidSocket{}
	return valid.Perform(ctx, ID, keyword, media, s.userProvider)
}

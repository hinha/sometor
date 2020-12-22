package infrastructure

import "github.com/hinha/sometor/provider"

type Infrastructure struct{}

// Fabricate infrastructure interface for kalkula
func Fabricate() (*Infrastructure, error) {
	i := &Infrastructure{}
	return i, nil
}

func (i *Infrastructure) FabricateCommand(cmd provider.Command) error {
	cmd.InjectCommand()

	return nil
}

func (i *Infrastructure) Close() {

}

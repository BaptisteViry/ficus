package plant

import (
	"context"
	"net/url"
	"trefle-provider/trefle/trefleutil"
)

type FakeClient struct {
	MockPlants []Plant

	UsedOpts url.Values
}

func (fc *FakeClient) List(context.Context, ...trefleutil.Option) ([]Plant, error) {
	return nil, nil
}

func (fc *FakeClient) Search(_ context.Context, query string, opts ...trefleutil.Option) ([]Plant, trefleutil.Meta, error) {
	var v = make(url.Values)

	for _, opt := range opts {
		opt.Set(v)
	}

	trefleutil.QueryOption(query).Set(v)

	fc.UsedOpts = v

	return fc.MockPlants, trefleutil.Meta{Total: int64(len(fc.MockPlants))}, nil
}

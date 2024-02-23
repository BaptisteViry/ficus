package plant

import (
	"context"
	"encoding/json"
	"net/http"

	"trefle-provider/trefle/trefleutil"
)

const entity = "plants"

type Client interface {
	Search(context.Context, string, ...trefleutil.Option) ([]Plant, trefleutil.Meta, error)
}

type client struct {
	*trefleutil.Client
}

func NewClient(cl *http.Client) Client {
	return &client{
		Client: &trefleutil.Client{
			Config: trefleutil.Config{
				Client: cl,
				Entity: entity,
			},
		},
	}
}

func (c *client) Search(ctx context.Context, query string, opts ...trefleutil.Option) ([]Plant, trefleutil.Meta, error) {
	if query == "" {
		return nil, trefleutil.Meta{}, nil
	}

	opts = append(opts, trefleutil.QueryOption(query))

	jsonPlants, err := c.Client.Search(ctx, opts...)

	if err != nil {
		return nil, trefleutil.Meta{}, nil
	}

	var rawPayload RawSearchPayload

	if err := json.Unmarshal(jsonPlants, &rawPayload); err != nil {
		return nil, trefleutil.Meta{}, nil
	}

	var meta = trefleutil.Meta{Total: rawPayload.Meta.Total}

	return rawPayload.Data, meta, nil
}

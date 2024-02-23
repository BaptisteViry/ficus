package plantmanager

import (
	"context"
	"net/http"

	"ficus/branches/nursery"
	"ficus/branches/nursery/provider"

	"trefle-provider/trefle/plant"
	"trefle-provider/trefle/trefleutil"
)

type Handler struct {
	provider.UnimplementedPlantManagerServer

	pcl plant.Client
}

func NewHandler(cl *http.Client) *Handler {
	return &Handler{pcl: plant.NewClient(cl)}
}

func (h *Handler) FetchPlants(ctx context.Context, req *provider.FetchPlantsRequest) (*provider.FetchPlantsResponse, error) {
	if req.Page == nil || req.Page.Size < 1 || req.Query == "" {
		return &provider.FetchPlantsResponse{}, nil
	}

	// Page.Size not used, trefle API does not allow to update it
	var opts = []trefleutil.Option{trefleutil.PageOption(req.Page.Number)}

	data, meta, err := h.pcl.Search(ctx, req.Query, opts...)

	if err != nil {
		return nil, err
	}

	var res = provider.FetchPlantsResponse{Total: int32(meta.Total)}

	for _, pl := range data {
		res.Plants = append(
			res.Plants,
			&nursery.Plant{
				Id:             pl.ID,
				CommonName:     pl.CommonName,
				ScientificName: pl.ScientificName,
				ImageUrl:       pl.ImageUrl,
			},
		)
	}

	return &res, nil
}

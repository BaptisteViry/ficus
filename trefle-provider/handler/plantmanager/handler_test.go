package plantmanager

import (
	"context"
	"net/url"
	"testing"

	"ficus/branches/nursery"
	"ficus/branches/nursery/provider"

	"github.com/stretchr/testify/assert"

	"trefle-provider/trefle/plant"
)

func TestFetchPlants(t *testing.T) {
	var (
		mockPlants = []plant.Plant{
			{
				ID:             1,
				CommonName:     "c1",
				ScientificName: "s1",
				ImageUrl:       "http://image_1.svg",
			},
			{
				ID:             2,
				CommonName:     "c2",
				ScientificName: "s2",
				ImageUrl:       "http://image_2.svg",
			},
		}

		nurseryPlants = []*nursery.Plant{
			{
				Id:             1,
				ScientificName: "s1",
				ImageUrl:       "http://image_1.svg",
				CommonName:     "c1",
			},
			{
				Id:             2,
				ScientificName: "s2",
				ImageUrl:       "http://image_2.svg",
				CommonName:     "c2",
			},
		}
	)

	for _, tt := range []struct {
		name        string
		query       string
		pageSize    int32
		wantOptions url.Values
		wantPlants  []*nursery.Plant
	}{
		{
			name:     "success",
			query:    "pachira aquatica",
			pageSize: 10,
			wantOptions: url.Values{
				"page": []string{"2"},
				"q":    []string{"pachira aquatica"},
			},
			wantPlants: nurseryPlants,
		},
		{name: "page size 0 no result"},
		{name: "no query no result"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var (
				fc = plant.FakeClient{MockPlants: mockPlants}
				h  = Handler{pcl: &fc}
			)

			res, err := h.FetchPlants(
				context.Background(),
				&provider.FetchPlantsRequest{
					Query: tt.query,
					Page:  &provider.Page{Number: 2, Size: tt.pageSize},
				},
			)

			assert.Nil(t, err)
			assert.Equal(t, tt.wantOptions, fc.UsedOpts)
			assert.Equal(t, tt.wantPlants, res.Plants)
			assert.Equal(t, int32(len(res.Plants)), res.Total)
		})
	}
}

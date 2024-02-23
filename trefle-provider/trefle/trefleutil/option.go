package trefleutil

import (
	"net/url"
	"strconv"
)

type Option interface {
	Set(url.Values)
}

type QueryOption string

func (qo QueryOption) Set(vs url.Values) {
	if qo == "" {
		return
	}

	vs.Add("q", string(qo))
}

type PageOption int64

func (po PageOption) Set(vs url.Values) {
	if po == 0 {
		return
	}

	vs.Add("page", strconv.Itoa(int(po)))
}

type PageSizeOption int64

func (pso PageSizeOption) Set(vs url.Values) {
	if pso == 0 {
		return
	}

	vs.Add("per_page", strconv.Itoa(int(pso)))
}

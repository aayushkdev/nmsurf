package ethernet

import "github.com/aayushkdev/nmsurf/internal/core"

type Provider struct{}

func New() *Provider {
	return &Provider{}
}

func (p *Provider) Type() core.NetworkType {
	return core.TypeEthernet
}

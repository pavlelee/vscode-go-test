package pkg

import (
	"context"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/server/mux"
	clusterprovider "tkestack.io/tke/pkg/platform/provider/cluster"
	"tkestack.io/tke/pkg/platform/types"
	v1 "tkestack.io/tke/pkg/platform/types/v1"
)

type Provider struct {
}

var _ clusterprovider.Provider = &Provider{}

func (p *Provider) Name() string {
	panic("not implemented") // TODO: Implement
}

func (p *Provider) RegisterHandler(mux *mux.PathRecorderMux) {
	panic("not implemented") // TODO: Implement
}

func (p *Provider) Validate(cluster *types.Cluster) field.ErrorList {
	panic("not implemented") // TODO: Implement
}

func (p *Provider) PreCreate(cluster *types.Cluster) error {
	panic("not implemented") // TODO: Implement
}

func (p *Provider) AfterCreate(cluster *types.Cluster) error {
	panic("not implemented") // TODO: Implement
}

// Setup called by controller to give an chance for plugin do some init work.
func (p *Provider) Setup() error {
	panic("not implemented") // TODO: Implement
}

// Teardown called by controller for plugin do some clean job.
func (p *Provider) Teardown() error {
	panic("not implemented") // TODO: Implement
}

func (p *Provider) OnCreate(ctx context.Context, cluster *v1.Cluster) error {
	panic("not implemented") // TODO: Implement
}

func (p *Provider) OnUpdate(ctx context.Context, cluster *v1.Cluster) error {
	panic("not implemented") // TODO: Implement
}

func (p *Provider) OnDelete(ctx context.Context, cluster *v1.Cluster) error {
	panic("not implemented") // TODO: Implement
}

// OnRunning call on first running.
func (p *Provider) OnRunning(ctx context.Context, cluster *v1.Cluster) error {
	panic("not implemented") // TODO: Implement
}

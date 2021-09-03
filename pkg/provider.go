package pkg

import (
	"github.com/pavlelee/vscode-test-module/pkg/provider"
	"github.com/pavlelee/vscode-test-module/pkg/registry"
)

type Provider struct {
}

var _ provider.Provider = &Provider{}

func (p *Provider) Name() string {
	registry.New(nil)
	panic("not implemented") // TODO: Implement
}

func (p *Provider) PreCreate() error {
	panic("not implemented") // TODO: Implement
}

package pkg

import (
	"github.com/pavlelee/vscode-test-module/pkg"
)

type Provider struct {
}

var _ pkg.Provider = &Provider{}

func (p *Provider) Name() string {
	panic("not implemented") // TODO: Implement
}

func (p *Provider) PreCreate() error {
	panic("not implemented") // TODO: Implement
}

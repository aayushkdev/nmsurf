package ethernet

import (
	"os/exec"

	"github.com/aayushkdev/nmsurf/internal/core"
)

func (p *Provider) Connect(n core.Network, password string) error {

	cmd := exec.Command(
		"nmcli",
		"device",
		"connect",
		n.Interface,
	)

	return cmd.Run()
}

func (p *Provider) Disconnect(n core.Network) error {

	cmd := exec.Command(
		"nmcli",
		"device",
		"disconnect",
		n.Interface,
	)

	return cmd.Run()
}

func (p *Provider) Forget(n core.Network) error {

	// Ethernet doesn't need forget
	return nil
}

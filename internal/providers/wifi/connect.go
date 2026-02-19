package wifi

import (
	"os/exec"

	"github.com/aayushkdev/nmsurf/internal/core"
)

func (p *Provider) Connect(n core.Network, password string) error {

	if password == "" {

		return exec.Command(
			"nmcli",
			"device",
			"wifi",
			"connect",
			n.BSSID,
		).Run()
	}

	return exec.Command(
		"nmcli",
		"device",
		"wifi",
		"connect",
		n.BSSID,
		"password",
		password,
	).Run()
}

func (p *Provider) Disconnect(n core.Network) error {

	return exec.Command(
		"nmcli",
		"connection",
		"down",
		n.BSSID,
	).Run()
}

func (p *Provider) Forget(n core.Network) error {

	return exec.Command(
		"nmcli",
		"connection",
		"delete",
		n.SSID,
	).Run()
}

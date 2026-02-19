package core

type Provider interface {
	Type() NetworkType

	Scan(hard bool) ([]Network, error)

	Connect(network Network, password string) error

	Disconnect(network Network) error

	Forget(network Network) error
}

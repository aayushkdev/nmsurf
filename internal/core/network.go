package core

type NetworkType int

const (
	TypeWiFi NetworkType = iota
	TypeEthernet
	TypeBluetooth
	TypeVPN
)

type Network struct {
	Type NetworkType

	// Wi-Fi specific
	SSID  string
	BSSID string

	// Ethernet specific
	Interface string

	// VPN specific
	UUID string

	Signal int

	Security string

	Frequency int
	Channel   int

	Connected bool
	Saved     bool
	Secured   bool
}

func (n Network) UniqueID() string {

	switch n.Type {

	case TypeWiFi:
		return n.BSSID

	case TypeEthernet:
		return n.Interface

	case TypeVPN:
		return n.UUID

	case TypeBluetooth:
		return n.Interface
	}

	return ""
}


func (n Network) DisplayName() string {

	switch n.Type {

	case TypeWiFi:
		return n.SSID

	case TypeEthernet:
		return n.Interface

	case TypeVPN:
		return n.UUID

	case TypeBluetooth:
		return n.Interface
	}

	return ""
}

func FreqToBand(freq int) string {

	switch {
	case freq >= 2400 && freq < 2500:
		return "2.4"

	case freq >= 4900 && freq < 5900:
		return "5"

	case freq >= 5900 && freq < 7125:
		return "6"

	default:
		return ""
	}
}

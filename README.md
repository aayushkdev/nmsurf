# nmsurf

A fast, minimal NetworkManager frontend for application launchers, written in Go.

nmsurf provides a simple interface to manage networks using NetworkManager, with support for launcher-based menus such as wofi and rofi.

---

## Screenshots


### Network List



### Network Menu



### Network Details



---

## Features

* Scan available Wi-Fi networks
* Connect, disconnect, and forget networks
* Hard rescan (hardware scan)
* View detailed network information
* Supports both rofi and wofi launchers
* Use your existing launcher theme or specify a custom theme
* Fast and lightweight

---

## Installation

### Prerequisites

* Wofi/rofi
* NetworkManager
* Go 1.20+ (build only/go install)


### Option 1: Arch User Repository (AUR)

For Arch Linux users, nmsurd is available in the AUR. This is the recommended method for easy installation and updates.

You can use any AUR helper like yay, paru, etc.

```bash
yay -S torrcli
```

---

### Option 2: Using Go

```bash
go install github.com/aayushkdev/nmsurf/cmd/nmsurf@latest
```

Ensure Go bin directory is in PATH:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

---

### Option 3: Using install.sh (works on any distro)

Building:

```bash
git clone https://github.com/aayushkdev/nmsurf.git
cd nmsurf
go build -ldflags="-s -w" -o nmsurf
```

Install:

```bash
./install.sh
```

This installs to:

```
/usr/local/bin/nmsurf
```

---

## Configuration

nmsurf can be configured using:

```bash
~/.config/nmsurf/config.toml
```

Example:

```toml
launcher = "wofi"
theme = "~/.config/wofi/style.css"
```

Or using rofi:

```toml
launcher = "rofi"
theme = "~/.config/rofi/network.rasi"
```

If no config file exists, defaults are used (wofi)

---

## Usage

Run:

```bash
nmsurf
```

Bind to key (Hyprland example):

```
bind = SUPER, N, exec, nmsurf
```

---

## License

MIT License


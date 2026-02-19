# nmsurf

A fast, minimal NetworkManager frontend for application launchers, written in Go.

nmsurf provides a simple interface to manage networks using NetworkManager, with support for launcher-based menus such as wofi.

---

## Screenshots


### Network List

<img width="1920" height="1080" alt="Screenshot-2026-02-19_23:50:54" src="https://github.com/user-attachments/assets/bb648e04-5941-40d5-af71-728e6aff4fe4" />

### Network Menu

<img width="1920" height="1080" alt="Screenshot-2026-02-19_23:51:17" src="https://github.com/user-attachments/assets/b8af8af8-b996-459c-8184-9fe667b63c3c" />


### Network Details

<img width="1920" height="1080" alt="Screenshot-2026-02-19_23:52:39" src="https://github.com/user-attachments/assets/81f700fd-9a9e-4a41-a96d-3f681bfbcf6f" />


---

## Features

* Scan available Wi-Fi networks
* Connect, disconnect, and forget networks
* Hard rescan (hardware scan)
* View detailed network information
* Supports:
    * Wofi
    * Rofi
    * Walker
    * Fuzzel
* Use your existing launcher theme or specify a custom theme
* Fast and lightweight

---

## Installation

### Prerequisites

* wofi, rofi, fuzzel or walker
* NetworkManager
* Go 1.20+ (build only/go install)


### Option 1: Arch User Repository (AUR)

For Arch Linux users, nmsurf is available in the AUR. This is the recommended method for easy installation and updates.

You can use any AUR helper like yay, paru, etc.

```bash
yay -S nmsurf
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
valid launchers include `wofi`, `rofi`, `walker`, `fuzzel`

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

wofi or any launcher based menu (except maybe rofi) does not expose any way to refresh the state of the menu without reloading it, so the ux may not be par with a tui or a gui based tool

---

## License

MIT License


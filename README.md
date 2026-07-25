# typr

A terminal-based typing test inspired by [toipe](https://github.com/Samyak2/toipe/).

<video src="assets/output.webm" autoplay loop muted></video>

## Usage

```sh
go run .
```

Type the displayed words and see your WPM and accuracy when finished.

- `ctrl-r` — restart
- `ctrl-c` — quit

## Build

```sh
make build       # current platform
make build-all   # linux, macOS, windows
```

Cross-platform — built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and pure Go.

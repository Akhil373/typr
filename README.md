# typr

A terminal-based typing test inspired by [toipe](https://github.com/Samyak2/toipe/).

[output.webm](https://github.com/user-attachments/assets/f180852e-9939-4c49-ba89-0a33997a5c5f)

Type the displayed words and see your WPM and accuracy when finished.
- run `typr -w <words>`
- `ctrl-r` — restart
- `ctrl-c` — quit

## Easy installation scripts

```ps1
curl https://raw.githubusercontent.com/Akhil373/typr/refs/heads/main/install.sh | bash # Linux or MacOS 
iwr https://raw.githubusercontent.com/Akhil373/typr/refs/heads/main/install.sh | iex # Windows
```

## ... or build it yourself

```sh
make build       # current platform
make build-all   # linux, macOS, windows
```

Cross-platform built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and pure Go.

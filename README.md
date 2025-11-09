# Maze Image CLI

A lightweight command-line tool that generates maze PNG images using the Hunt-and-Kill algorithm. Width, height, cell size, and output path are all configurable flags so you can quickly create printable mazes without running any additional UI or viewer.

## Features
- Hunt-and-Kill maze generation implemented in `lib/maze.go`.
- Single CLI (`cmd/mazeimg`) that both generates the maze and renders it as a PNG.
- Configurable image size through `-width`, `-height`, and `-cell` flags.

## Requirements
- Go 1.20 or newer

## Getting Started
```sh
git clone https://github.com/storbeck/maze.git
cd maze
```

### Build once
```sh
go build -o bin/mazeimg ./cmd/mazeimg
```

### Or run directly
```sh
go run ./cmd/mazeimg -width 20 -height 20 -cell 16 -out maze.png
```

### CLI Flags
| Flag | Default | Description |
| --- | --- | --- |
| `-width` | 20 | Number of maze cells horizontally |
| `-height` | 20 | Number of maze cells vertically |
| `-cell` | 16 | Pixel size of each rune in the rendered grid |
| `-out` | `maze.png` | Output file path |
| `-letter` | `false` | Auto-size the PNG for US Letter (cell size ignored) |

The CLI prints the destination file once the render completes. Larger mazes or bigger cell sizes produce proportionally larger PNGs.

Use `-letter` to generate a maze pre-sized for printing on US Letter paper at 300 dpi (roughly 7.4"×9.4" of maze). The CLI automatically rescales the cell size to keep the image within printable bounds, so you can still raise `-width`/`-height` for added complexity without worrying about manual math. You can override `-out`, but `-cell` is ignored whenever `-letter` is used.

## Directory Structure
```
.
├── cmd/mazeimg        # CLI entry point
├── go.mod
├── go.sum
├── lib                # Maze generation logic
└── README.md
```

## Cleaning Up
Remove the built binary and generated image files as needed:
```sh
rm -f bin/mazeimg *.png
```

# Maze Generator and 2D Viewer

This project provides tools to generate a maze and view it in a simple 2D ASCII representation. 

## Repository Structure

```
.
├── go.mod
├── lib
│   └── maze.go
├── gen
│   └── main.go
└── 2d
    └── main.go
```

## Prerequisites

- Go 1.16+ installed on your machine.

## Installation

Clone the repository:

```sh
git clone https://github.com/storbeck/maze.git
cd maze
```

Initialize the Go module:

```sh
go mod init github.com/storbeck/maze
```

## Building

Build both the maze generator and the 2D viewer from the root directory:

```sh
go build -o mazegen gen/main.go
go build -o maze2d 2d/main.go
```

## Usage

### Generating a Maze

Generate a maze using the `mazegen` tool. Specify the size using the `-size` flag (e.g., 10x10):

```sh
./mazegen -size 10x10 > maze.txt
```

### Viewing the Maze in 2D

View the generated maze in a 2D ASCII representation using the `maze2d` tool:

```sh
./maze2d maze.txt
```

## Example

Generate and view a 10x10 maze:

```sh
./mazegen -size 10x10 > maze.txt
./maze2d maze.txt
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
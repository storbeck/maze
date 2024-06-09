# Maze Generator and 3D Viewer

This project provides tools to generate a maze and view it in a simple 3D-like ASCII representation. The project is organized into three main components:
1. **mazelib**: A library for maze generation and representation.
2. **mazegen**: A command-line tool to generate mazes.
3. **maze3d**: A command-line tool to display the maze in a 3D-like ASCII representation.

## Repository Structure

```
.
├── go.mod
├── lib
│   └── maze.go
├── gen
│   └── main.go
└── 3d
    └── main.go
```

## Prerequisites

- Go 1.16+ installed on your machine.
- Git installed on your machine.

## Installation

First, clone the repository:

```sh
git clone https://github.com/storbeck/maze.git
cd maze
```

Initialize the Go module (if not already done):

```sh
go mod init github.com/storbeck/maze
```

## Building

You can build both the maze generator and the 3D viewer from the root directory.

### Building the Maze Generator

```sh
go build -o mazegen gen/main.go
```

### Building the 3D Maze Viewer

```sh
go build -o maze3d 3d/main.go
```

## Usage

### Generating a Maze

To generate a maze, use the `mazegen` tool. You can specify the size of the maze using the `-size` flag. For example, to generate a 10x10 maze:

```sh
./mazegen -size 10x10 > maze.txt
```

This will create a 10x10 maze and output it to `maze.txt`.

### Viewing the Maze in 3D

To view the generated maze in a 3D-like ASCII representation, use the `maze3d` tool:

```sh
./maze3d maze.txt
```

This will read the maze from `maze.txt` and display it in the console.

## Example

Here is an example of generating and viewing a 10x10 maze:

1. **Generate the Maze:**

   ```sh
   ./mazegen -size 10x10 > maze.txt
   ```

2. **View the Maze in 3D:**

   ```sh
   ./maze3d maze.txt
   ```

## Contributing

Feel free to submit issues or pull requests if you have any improvements or bug fixes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```

### Additional Steps
Ensure your `go.mod` file is properly configured for the module:
```sh
module github.com/storbeck/maze

go 1.16
```

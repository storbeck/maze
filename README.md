# Maze Generator and Viewer



https://github.com/storbeck/maze/assets/449874/a9d34641-b42e-44bc-8a47-9cc261acef9f



This project provides tools to generate and view mazes in both ASCII and 2D formats. The mazes are generated using a challenging Hunt-and-Kill algorithm to ensure a complex and engaging puzzle experience.

## Features

- **Maze Generation**: Generate complex mazes using the Hunt-and-Kill algorithm.
- **ASCII Viewer**: View the generated mazes in a simple ASCII format.
- **2D Viewer**: View the generated mazes in a top-down 2D style with textures.

## Prerequisites

- Go 1.16 or higher

## Directory Structure

```
.
├── 2d
│   └── maze.go            # Main program for generating and viewing mazes in 2D format with textures
├── ascii
│   └── maze.go            # Main program for generating and viewing mazes in ASCII format
├── gen
│   └── maze.go            # Maze generation logic using the Hunt-and-Kill algorithm
├── lib
│   └── maze.go            # Shared library for maze generation algorithms
├── scripts
│   ├── 2d.sh              # Script to generate a new 2D maze and view it
│   └── ascii.sh           # Script to generate a new ASCII maze and view it
├── textures
│   ├── entrance
│   │   └── generate.go    # Generator for the entrance texture
│   ├── exit
│   │   └── generate.go    # Generator for the exit texture
│   ├── floor
│   │   └── generate.go    # Generator for the floor texture
│   └── wall
│       └── generate.go    # Generator for the wall texture
├── bin                    # Directory for compiled binaries
├── Makefile               # Makefile for building the project and generating textures
├── README.md              # This README file
├── go.mod                 # Go module file
└── go.sum                 # Go dependencies file
```

## Installation

1. **Clone the repository**:
   ```sh
   git clone https://github.com/storbeck/maze.git
   cd maze
   ```

2. **Build and generate textures**:
   ```sh
   make
   ```

## Usage

### Generate and View a New 2D Maze

To generate a new maze, generate textures, and view it in 2D:

```sh
make && ./scripts/2d.sh
```

### Generate and View a New ASCII Maze

To generate a new maze and view it in ASCII without regenerating textures:

```sh
make build && ./scripts/ascii.sh
```

## Generating Textures Only

If you need to regenerate the textures separately, run:

```sh
make textures
```

## Cleaning Up

To clean up the generated files, run:

```sh
make clean
```

# Maze Generator

This project generates mazes and displays them in 2D and 3D with various textures. 

![maze](https://github.com/storbeck/maze/assets/449874/bc4cb407-ea7f-4f25-9121-fc3de7ec8a54)

## Setup

Ensure you have [Go](https://golang.org/dl/) installed on your system.

### Usage

1. **Clone the repository**:
   ```sh
   git clone https://github.com/storbeck/maze.git
   cd maze
   ```

2. **Generate the textures and build the project**:
   ```sh
   make && ./new-3d.sh
   ```
   This command will:
   - Generate new textures.
   - Build the `mazegen`, `maze2d`, and `maze3d` executables.
   - Generate a new 3D maze image and open it.

3. **Build the project without generating new textures for 2D**:
   ```sh
   make build && ./new-2d.sh
   ```
   This command will:
   - Build the `mazegen`, `maze2d`, and `maze3d` executables without generating new textures.
   - Generate a new 2D maze image and open it.

### Makefile Commands

- `make`: Generates new textures and builds the project.
- `make build`: Builds the project without generating new textures.
- `make clean`: Cleans up the generated executables and textures.

### Texture Generators

- **Floor Texture**: Generates a tiled floor texture.
- **Wall Texture**: Generates a gradient wall texture.
- **Lava Texture**: Generates a noisy lava-like texture.
- **Polka Dot Texture**: Generates a polka dot pattern texture.

### Shell Scripts

- **`new-3d.sh`**: Generates a new maze, creates new textures, generates a 3D image, and opens it.
- **`new-2d.sh`**: Generates a new maze and generates a 2D image without creating new textures.

### Dependencies

- [Go](https://golang.org/dl/)
- [gg](https://github.com/fogleman/gg) library for graphics:
  ```sh
  go get -u github.com/fogleman/gg
  ```

### Example

To generate and view a new 3D maze with textures:
```sh
make && ./new-3d.sh
```

To generate and view a new 2D maze without generating new textures:
```sh
make build && ./new-2d.sh
```

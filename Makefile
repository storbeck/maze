.PHONY: all textures build clean

all: textures build

textures:
	go run textures/floor/generate.go
	go run textures/wall/generate.go
	go run textures/entrance/generate.go
	go run textures/exit/generate.go

build:
	go build -o mazegen gen/maze.go
	go build -o maze2d 2d/maze.go
	go build -o maze3d 3d/maze.go

clean:
	rm -f mazegen maze2d maze3d
	rm -f textures/*.png

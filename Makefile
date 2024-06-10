.PHONY: all textures build clean

all: textures build

textures:
	go run textures/floor/generate.go
	go run textures/wall/generate.go
	go run textures/entrance/generate.go
	go run textures/exit/generate.go

build:
	go build -o bin/mazegen gen/maze.go
	go build -o bin/mazeascii ascii/maze.go
	go build -o bin/maze2d 2d/maze.go

clean:
	rm -f bin/mazegen bin/mazeascii bin/maze2d
	rm -f textures/*.png
	rm -f maze.txt maze.png

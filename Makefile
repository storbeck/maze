.PHONY: all textures build clean

all: textures build

textures:
	cd textures && go run brick/generate.go
	cd textures && go run grass/generate.go
	cd textures && go run wall/generate.go
	cd textures && go run floor/generate.go
	cd textures && go run lava/generate.go
	cd textures && go run polkadot/generate.go

build:
	go build -o mazegen gen/maze.go
	go build -o maze2d 2d/maze.go
	go build -o maze3d 3d/maze.go

clean:
	rm -f mazegen maze2d maze3d


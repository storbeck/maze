.PHONY: build run clean

build:
	mkdir -p bin
	go build -o bin/mazeimg ./cmd/mazeimg

run:
	go run ./cmd/mazeimg $(ARGS)

clean:
	rm -f bin/mazeimg *.png

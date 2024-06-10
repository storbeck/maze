#!/bin/bash

SIZE=$1

if [ -z "$SIZE" ]; then
    SIZE=10x10
fi

./bin/mazegen -size $SIZE > maze.txt
./bin/maze2d maze.txt maze.png
imgcat maze.png

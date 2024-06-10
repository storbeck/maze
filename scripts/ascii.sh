#!/bin/bash

SIZE=$1

if [ -z "$SIZE" ]; then
    SIZE=10x10
fi

./mazegen -size $SIZE > maze.txt
./maze2d maze.txt

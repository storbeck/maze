#!/bin/bash

./mazegen -size 10x10 > maze.txt
./maze3d maze.txt maze.png
open maze.png

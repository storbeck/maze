#!/bin/bash

./mazegen -size 20x10 > maze.txt
./maze3d maze.txt maze.png
open maze.png

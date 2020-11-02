# Advent of Code 2019

This repository contains Go solutions and helper libraries for Advent of Code 2019.

## Setup
All executables contained within this repo can be built with a single `make` invocation:
```bash
$ make build
```
Resultant executables will be located in `./bin`.

## Execution
All solutions are executed directly from the command line.  If a problem requires an input, it is accepted from STDIN:
```bash
$ ./cmd/day11/input.txt | ./bin/day11
```
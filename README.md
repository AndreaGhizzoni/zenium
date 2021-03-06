# Project Zenium
[![codebeat badge](https://codebeat.co/badges/42816dbb-cc47-4931-9685-39d7b5ab2281)](https://codebeat.co/projects/github-com-andreaghizzoni-zenium-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/AndreaGhizzoni/zenium)](https://goreportcard.com/report/github.com/AndreaGhizzoni/zenium)

Zenium it'a a utility program to generate random data files.

## Introductions
* [Dependencies](#dependencies)
* [Installation](#installation)
* [Structure of Data](#structure-of-generated-data)
    - [Random Slice](#slice)
    - [Random Matrix](#matrix)
    - [Random Bound](#bound)
* [TODO](#todo)
    
## Dependencies
 - [cli framework](https://github.com/urfave/cli)

## Installation
Keep in mind that this project is still WORK IN PROGRESS, so changes will be
made.
```bash
cd $GOPATH/src
go get github.com/AndreaGhizzoni/zenium
./$GOPATH/bin/zenium -h
```

## Structure of Generated Data
All generated data will be in the following format:

### Slice
```bash
./$GOPATH/bin/zenium rslice -l 50 -o vector.txt -m 1 -M 10
```
This command generate a file called `vector.txt` that contains `50` random 
numbers, from `1` to `10` using current time as seed to generate it.

The file format is the follow:
```
n
x1 x2 [...] xn
```
Where `n` is the given length of vector (50 in the example) and `x1 x2 [...] xn`
are all the generated numbers (with the properties see previously in the 
example) separated by space.


### Matrix
```bash
./$GOPATH/bin/zenium matrix -o matrix.txt -c 10 -r 10 -m 1 -M 10
```
This command generate a file called `matrix.txt` that contains a `10x10` matrix
of random numbers, from `1` to `10` using current time as seed to generate it.

The file format is the follow:
```
m n
x11 [...] x1n
x21
[...]
xm1 [...] xmn
```
Where `m n` is the given size of matrix (10 10 in the example) and listed below
that all the elements of matrix: columns separated by space and rows by new line.


### Bound
```bash
./$GOPATH/bin/zenium bound -o bound.txt -m -10 -M 10 -w 3 -a 10
```
This command generate a file called `bound.txt` that contains a list of 
`10` bounds of width `3`, all in between `-10` and `10`.

The file format is the follow:
```
a
x1 y1
x2 y2
[...]
xa ya
```
Where `a` is the amount of bounds listed below, and for each line the bound 
is expressed as lower bound `x` and upper bound `y`, where `m <= x < y <= M`. 

# TODO
- [x] Slice (aka Vector)
- [x] Matrix
- [x] Complete -output|-o flag
- [x] Bound
- [x] Using math/big instead of all int64
- [x] Using crypto/rand in certain cases
- [ ] Ordered Slice
- [ ] Multiple generation with a single command
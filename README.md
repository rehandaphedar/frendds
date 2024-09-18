# Introduction

![frendds logo](https://git.sr.ht/~rehandaphedar/frendds/blob/main/favicon.png)

frendds is a program to generate a [D2 file](https://d2lang.com) from [friends.txt](https://sr.ht/~rehandaphedar/friends.txt) files of a domain.

# Examples

An example output generated by `frendds rehandaphedar.com` is in [example.d2](https://git.sr.ht/~rehandaphedar/frendds/tree/main/item/example.d2). The output generated by `d2 example.d2` with default settings looks like this:

![example.svg](https://git.sr.ht/~rehandaphedar/frendds/blob/main/example.svg)

# Installation

## Install Dependencies

- `go`

## Installation

```shell
go install git.sr.ht/~rehandaphedar/frendds@latest
```

# Usage

```shell
frendds [domain]
```

Example:

```shell
frendds rehandaphedar.com
```

You can use the generated file with `d2`. For example, to generate an svg:

```shell
frendds rehandaphedar.com | d2 - graph.svg
```

The output is a normal `d2` file, and thus, all options from the [D2 documentation](https://d2lang.com) should work.

# Building

## Build Dependencies

- `go`

## Building

Clone the source code

```shell
git clone https://git.sr.ht/~rehandaphedar/frendds
cd frendds
go build .
```

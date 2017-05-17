teets
=========

[![Build Status](https://travis-ci.org/servak/teets.svg?branch=master)](https://travis-ci.org/servak/teets)

Tee and Timestamp in Golang.

## Usage

```
> ./teets
Usage:
  ./teets [options] filepath

Options:
  -a    append flag
  -f string
        timestamp format. https://golang.org/src/time/format.go (default "15:04:05")
  -o    add timestamp in stdout

> echo foooo | ./teets aa.log
foooo
> cat aa.log
[22:57:38] foooo
[22:57:38] 
```

# Go example project

[![Build Status](https://badge.buildkite.com/399f24bbeddf266368de484129387bb437110cf475c11708e1.svg)](https://buildkite.com/lelia/buildkite-pipeline-example) [![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)

This repository is a fork of the basic Golang example repo, trimmed down to contain a single example.

## Build the project

```
$ cd hello
$ go build
```

A simple application that takes a command line argument, and then returns it to you in a string:

```
$ chmod +x hello/hello
$ ./hello/hello John Doe
```

The above will return 'Hello, John Doe!'

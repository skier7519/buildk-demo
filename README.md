# Go example project

[![Build status](https://badge.buildkite.com/a2febd9a702fdaeeee0512ab36bb69c922871f1ee736d87a55.svg)](https://buildkite.com/lelia/buildkite-pipeline-example) [![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)

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

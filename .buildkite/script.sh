#!/bin/bash

name=$(buildkite-agent meta-data get name)

#echo "pwd"
#pwd
#echo "cd hello"
cd hello
#echo "pwd"
#pwd
#echo "go build"
go build
#echo "chmod+x hello"
chmod +x hello
#echo "./hello $name"
./hello $name

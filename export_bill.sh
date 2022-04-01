#!/bin/bash

go_file="go1.12.17.linux-amd64.tar.gz"

# wget https://go.dev/dl/${go_file}

# tar -xzf ${go_file}

ROOT_DIR=`pwd`

export PATH=${ROOT_DIR}/go/bin/:$PATH

go version

go run ./main.go

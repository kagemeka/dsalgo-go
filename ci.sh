#!/bin/bash
apt update
apt install -y python3-pip
pip install -U pip
pip install pre-commit

go fmt -x ./...
pre-commit run --all-files

go test -v ./...

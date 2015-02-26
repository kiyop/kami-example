#!/bin/sh

CWD=$(cd $(dirname $0);pwd)
export GOPATH="${CWD}/_vendor"

dev_appserver.py app.yaml

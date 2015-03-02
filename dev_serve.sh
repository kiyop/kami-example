#!/bin/sh
CWD=$(cd $(dirname $0);pwd)
cd "${CWD}"

export GOPATH="${CWD}/_vendor"
#echo $GOPATH

dev_appserver.py \
	app.yaml \
	--host 0.0.0.0 \
	--enable_sendmail


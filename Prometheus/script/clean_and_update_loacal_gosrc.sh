#!/bin/bash
source /etc/profile
rm -rf ${GOPATH}/src/github.com/luckykris/Cronus/Prometheus
mkdir -p ${GOPATH}/src/github.com/luckykris/Cronus/Prometheus
cp -R ./* ${GOPATH}/src/github.com/luckykris/Cronus/Prometheus/

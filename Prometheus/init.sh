#!/usr/bin/env bash 
CURDIR=`pwd` 
if [ ! -d src ];then
mkdir $CURDIR/src
fi
if [ ! -d bin ];then
mkdir $CURDIR/bin
fi
if [ ! -d pkg ];then
mkdir $CURDIR/pkg
fi
export GOPATH="${CURDIR}/" 
echo 'finished' 


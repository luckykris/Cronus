#!/usr/bin/env bash 
CURDIR=/tmp 
if [ ! -d $CURDIR/src ];then
mkdir $CURDIR/src
fi
if [ ! -d $CURDIR/bin ];then
mkdir $CURDIR/bin
fi
if [ ! -d $CURDIR/pkg ];then
mkdir $CURDIR/pkg
fi
export GOPATH="${CURDIR}/" 
echo 'finished' 


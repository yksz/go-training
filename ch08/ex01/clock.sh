#!/bin/sh

cd `dirname $0`

go build clock.go
if [ $? -ne 0 ] ; then
    exit 1
fi

TZ=US/Eastern    ./clock -port 8010 &
TZ=Asia/Tokyo    ./clock -port 8020 &
TZ=Europe/London ./clock -port 8030 &

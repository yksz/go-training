#!/bin/sh

cd `dirname $0`
mkdir -p index
cd index
curl https://xkcd.com/[570-579]/info.0.json -o "#1.json"

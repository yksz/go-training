#!/bin/sh

cd `dirname $0`

go run clockwall.go NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030

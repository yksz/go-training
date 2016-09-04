#!/bin/bash

go build search.go
echo "[INFO] Start a test server"
./search &
pid=$!
sleep ${1:-1}

check() {
    if [ "$2" != "$3" ] ; then
        echo "$1 FAIL: got:$2, want:$3"
        return 1
    fi
    echo "$1 OK"
}

got1=`fetch 'http://localhost:12345/search'`
want1='Search: {Labels:[] MaxResults:10 Exact:false Email:}'
check 'Test1' "$got1" "$want1"

got2=`fetch 'http://localhost:12345/search?email=abc@example.com'`
want2='Search: {Labels:[] MaxResults:10 Exact:false Email:abc@example.com}'
check 'Test2' "$got2" "$want2"

got3=`fetch 'http://localhost:12345/search?email=abc'`
want3="'email' is not email"
check 'Test3' "$got3" "$want3"

echo "[INFO] Stop the test server"
kill $pid

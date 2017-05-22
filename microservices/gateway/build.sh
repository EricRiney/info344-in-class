#!/usr/bin/env bash
if [ -z "$1" ]
then
    echo 'usage:'
    echo '  ./build.sh <container-image-tag>'
    exit 1
fi

echo "building Linux executable..."
GOOS=linux go build
echo "building Docker container image..."
docker build -t $1 .
echo "cleaning up..."
go clean 

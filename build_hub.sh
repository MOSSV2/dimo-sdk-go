#!/bin/sh

version=$1

docker build -t mossv2/dimo:hub-$version -f Dockerfile .
docker tag mossv2/dimo:hub-$version mossv2/dimo:hub-latest
docker push mossv2/dimo:hub-$version
docker push mossv2/dimo:hub-latest

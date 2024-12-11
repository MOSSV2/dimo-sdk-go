#!/bin/sh

version=$1

docker build -t dimo-hub:$version -f Dockerfile .
docker tag dimo-hub:$version dimo-hub:latest
docker tag dimo-hub:$version mossv2/dimo:hub-$version
docker tag dimo-hub:$version mossv2/dimo:hub-latest
docker push mossv2/dimo:hub-$version
docker push mossv2/dimo:hub-latest

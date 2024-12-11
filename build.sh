#!/bin/sh

version=$1

docker build --build-arg DIMO_BIN=hub-edge -t dimo-hub:$version -f Dockerfile.local .
docker tag dimo-hub:$version dimo-hub:latest
docker tag dimo-hub:$version mossv2/dimo:hub-$version
docker tag dimo-hub:$version mossv2/dimo:hub-latest
docker push mossv2/dimo:hub-$version
docker push mossv2/dimo:hub-latest

cp ../dimo-go/bin/store-edge ./bin
docker build --build-arg DIMO_BIN=store-edge -t dimo-store:$version -f Dockerfile.local .
docker tag dimo-store:$version dimo-store:latest
docker tag dimo-store:$version mossv2/dimo:store-$version
docker tag dimo-store:$version mossv2/dimo:store-latest
docker push mossv2/dimo:store-$version
docker push mossv2/dimo:store-latest

cp ../dimo-go/bin/stream-edge ./bin
docker build --build-arg DIMO_BIN=stream-edge -t dimo-stream:$version -f Dockerfile.local .
docker tag dimo-stream:$version dimo-stream:latest
docker tag dimo-stream:$version mossv2/dimo:stream-$verison
docker tag dimo-stream:$version mossv2/dimo:stream-latest
docker push mossv2/dimo:stream-$version
docker push mossv2/dimo:stream-latest

cp ../dimo-go/bin/validator-edge ./bin
docker build --build-arg DIMO_BIN=validator-edge -t dimo-validator:$version -f Dockerfile.local .
docker tag dimo-validator:$version dimo-validator:latest
docker tag dimo-validator:$version mossv2/dimo:validator-$version
docker tag dimo-validator:$version mossv2/dimo:validator-latest
docker push mossv2/dimo:validator-$version
docker push mossv2/dimo:validator-latest

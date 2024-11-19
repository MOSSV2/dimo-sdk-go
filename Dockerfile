# Dockerfile for development purposes.

###############################################################################
# Base build image
FROM golang:1.22-alpine AS build_base
RUN apk add bash ca-certificates git gcc g++ libc-dev

ENV GO111MODULE=on
# Populate the module cache based on the go.{mod,sum} files.
COPY . /opt/dimo
WORKDIR /opt/dimo

###############################################################################
FROM build_base AS server_builder
ARG TARGETARCH
ARG GIT_BRANCH="unknown"
ARG GIT_REVISION="unknown"
ARG BUILD_USER="unknown"
ARG BUILD_DATE="unknown"
ARG EXTRA_BUILD_ARGS=""
COPY . .
RUN GOOS=linux GOARCH=$TARGETARCH go build $EXTRA_BUILD_ARGS \
      -ldflags '-w -extldflags "-static"' \
      -o /hub ./app/hub

################################################################################
FROM alpine AS hub
ENTRYPOINT ["/bin/hub", "daemon", "run"]
COPY --from=server_builder /hub /bin/hub
RUN apk add --no-cache --upgrade bc ca-certificates openssl
CMD ["--bind", "0.0.0.0:8086"]
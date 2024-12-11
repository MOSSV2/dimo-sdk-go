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
COPY scripts/start.sh /app/start
RUN GOOS=linux GOARCH=$TARGETARCH go build $EXTRA_BUILD_ARGS \
      -ldflags '-w -extldflags "-static"' \
      -o /app/hub ./app/hub

################################################################################
FROM alpine AS hub

WORKDIR /app
ENTRYPOINT ["/app/start"]
COPY --from=server_builder /app/start /app/start
COPY --from=server_builder /app/hub /app/dimo
COPY --from=server_builder /opt/dimo/app/hub/assets /app/assets
#RUN apk add --no-cache --upgrade bc ca-certificates openssl
#CMD ["--bind", "0.0.0.0:8080"]
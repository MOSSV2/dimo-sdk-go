#!/bin/sh

REPO_DIR=/data


if [ ! -f $REPO_DIR/config.json ];then
  echo "===== need init ====="
  /app/dimo --repo=$REPO_DIR init
fi

echo "===== start ====="
exec /app/dimo --repo=$REPO_DIR daemon run --bind 0.0.0.0:8080 $@



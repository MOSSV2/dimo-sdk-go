#!/bin/sh

HOME_DIR=$HOME/.hub

echo $HOME_DIR


if [ ! -f $HOME_DIR/config.json ];then
  echo "===== need init ====="
  /bin/hub init
fi

echo "===== start ====="
/bin/hub daemon run --bind 0.0.0.0:8080




#! /bin/bash

./build/bin/geth --exec "$2" attach ipc:storage/node$1/geth.ipc


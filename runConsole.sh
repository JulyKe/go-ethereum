#! /bin/bash
./build/bin/geth --exec "$2" attach ipc:node$1/geth.ipc

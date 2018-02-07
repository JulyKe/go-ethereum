#! /bin/bash

./build/bin/geth --exec 'miner.stop()' attach ipc:storage/node$1/geth.ipc


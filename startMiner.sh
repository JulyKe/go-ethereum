#! /bin/bash

./build/bin/geth --exec 'miner.start()' attach ipc:storage/node$1/geth.ipc


#! /bin/bash
./build/bin/geth --exec 'miner.start(1)' attach ipc:storage/node$1/geth.ipc

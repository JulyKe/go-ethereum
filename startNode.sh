#! /bin/bash
./build/bin/geth -identity "node$1" --rpc --rpcport 854$1 --datadir storage/node$1 --port 3030$1 --nodiscover --networkid 1900 $2 2>>node$1.log

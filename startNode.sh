#! /bin/bash
 ./build/bin/geth --rpc --rpcport 854$1 --datadir node$1/ --port 3030$1 --nodiscover --networkid 1900 console

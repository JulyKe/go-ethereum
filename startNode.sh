#! /bin/bash
 #./build/bin/geth --rpc --rpcport 854$1 --datadir node$1/ --port 3030$1 --nodiscover --networkid 1900 console
 #./build/bin/geth -identity "node$1" --genesis genesis.json --rpc --rpcport 854$1 --datadir node$1/ --port 3030$1 --nodiscover --networkid 1900 $2 2>> node1Log &
 #./build/bin/geth -identity "node$1" --genesis genesis.json --rpc --rpcport 854$1 --datadir node$1/ --port 3030$1 --nodiscover --networkid 1900 $2 2>>log_node$1
 ./build/bin/geth -identity "node$1" --genesis genesis.json --datadir storage/node$1/ --port 3030$1 --nodiscover --networkid 1900 $2 2>>log_node$1

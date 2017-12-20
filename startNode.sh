#! /bin/bash
./build/bin/geth -identity "node$1" --genesis genesis.json --datadir storage/node$1/ --port 3030$1 --nodiscover --networkid 1900 $2 2>>log_node$1

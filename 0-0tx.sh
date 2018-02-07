#! /bin/bash
./build/bin/geth --exec 'personal.unlockAccount(eth.accounts[0],"account0")' attach ipc:storage/node$1/geth.ipc
./build/bin/geth --exec 'eth.sendTransaction({from:eth.accounts[0], to:eth.accounts[1], value:1})' attach ipc:storage/node$1/geth.ipc

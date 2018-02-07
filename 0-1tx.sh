#! /bin/bash
./build/bin/geth --exec 'personal.unlockAccount(eth.accounts[0],"account0")' attach ipc:storage/node$1/geth.ipc
./build/bin/geth --exec 'eth.sendTransaction({from:eth.accounts[0], to:"aa01810cef930c3e563e23d93256b27bfc4a73ee", value:1})' attach ipc:storage/node$1/geth.ipc

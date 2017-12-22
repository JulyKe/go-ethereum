#! /bin/bash
PID=`ps aux | grep geth | grep "./build/bin/" | grep node$1 | awk '{print $2}'`
echo $PID
kill $PID

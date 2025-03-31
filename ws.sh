#!/bin/bash
echo "***1**which go:"
which go
source ~/.zshrc
export PATH=$PATH:/usr/local/go/bin
echo "***2**which go:"
which go
echo "*****path info:" $PATH
go version
go mod download
go build -o wsapp ./ws.go
echo "current pwd:$(pwd)"
nohup ./wsapp < /dev/null > wsapp.log 2>&1 &
disown $!
ps aux | grep wsapp
sleep 5
echo "---------"
whoami
ps aux | grep wsapp
curl http://localhost:9002/json | jq
#exit 0
echo "-----2-----"
curl http://localhost:9002/json | jq

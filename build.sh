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
go build -o myapp .
echo "current pwd:$(pwd)"
nohup ./myapp < /dev/null > app.log 2>&1 &
disown $!
ps aux | grep myapp
sleep 5
echo "---------"
whoami
ps aux | grep myapp
curl http://localhost:9003/hello | jq
#exit 0
echo "-----2-----"
curl http://localhost:9003/hello | jq

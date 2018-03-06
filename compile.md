go get   -ldflags " -linkmode external -extldflags '-static -lpthread' "    github.com/hyperledger/fabric-ca/cmd/...

go get   -ldflags " -linkmode external -extldflags '-static -lpthread' "    github.com/hyperledger/fabric/orderer/...


fabric-ca-server init -b adf:adfasdfasdf -d


dlv debug  github.com/hyperledger/fabric-ca/cmd/fabric-ca-server/main --wd /tmp/ca -- init -b adf:adfasdfasdf -d

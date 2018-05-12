
#!/bin/bash
my_dir=`dirname $0`

protoc -I=$my_dir -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf --gogo_out=plugins=grpc:$my_dir/ $my_dir/*.proto
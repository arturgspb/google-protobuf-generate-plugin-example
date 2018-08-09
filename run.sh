#protoc --plugin=protoc-gen-demo=gen-impl.py --demo_out=. demo.proto
#protoc --python_out=./out --mypy_out=./out ./defs/demo.proto

protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --python_out=./out \
  --mypy_out=./out \
  --swagger_out=logtostderr=true:./out \
  --grpc-gateway_out=logtostderr=true:./out \
  --go_out=plugins=grpc:./out \
  ./defs/demo.proto


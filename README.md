Inspored by:
https://medium.com/swlh/building-a-realtime-dashboard-with-reactjs-go-grpc-and-envoy-7be155dfabfb

cd protos
# SERVER

protoc --go-grpc_out=server protos/user.proto && protoc --go_out=server protos/user.proto

# CLIENT

PROTOC_GEN_TS_PATH="../client/node_modules/.bin/protoc-gen-ts"
OUT_DIR="../client/src/pbs/user"
protoc --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" --js_out="import_style=commonjs,binary:${OUT_DIR}" --ts_out="${OUT_DIR}" user.proto 
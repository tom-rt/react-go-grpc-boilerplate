Inspired by:
https://medium.com/swlh/building-a-realtime-dashboard-with-reactjs-go-grpc-and-envoy-7be155dfabfb

# PROXY
docker build -t grpc-envoy:1.0 .
docker run --network=host grpc-envoy:1.0
# SERVER
go get .
go run server.go -port=8080
# CLIENT
npx create-react-app client --template typescript
cd client
npm install ts-protoc-gen
npm install grpc-web
npm install google-protobuf
npm install @improbable-eng/grpc-web

Add `
    "ignorePatterns": [
      "**/*_pb.js"
    ]
`
To "eslintConfig" in your package.json

npm start
# PROTOS

cd protos
## Client:
PROTOC_GEN_TS_PATH="../client/node_modules/.bin/protoc-gen-ts"
OUT_DIR="../client/src/pbs"
### models:
protoc --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" --js_out="import_style=commonjs,binary:${OUT_DIR}" --ts_out="${OUT_DIR}" base.proto 

### services:
protoc --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" --js_out="import_style=commonjs,binary:${OUT_DIR}" --ts_out="service=grpc-web:${OUT_DIR}" base.proto
## Server:
protoc --go-grpc_out=../server base.proto
protoc --go_out=../server base.proto


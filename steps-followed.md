
./bin/cryptogen generate --config=./artifacts/channel/cryptogen.yaml

export FABRIC_CFG_PATH=$PWD
./bin/configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./channel/genesis.block


Creating Chanel
export CHANNEL_NAME=mychannel


../bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel/mychannel.tx -channelID $CHANNEL_NAME


../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel/AgentMSPanchors.tx -channelID $CHANNEL_NAME -asOrg AgentMSP


../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel/IllinoisMSPanchors.tx -channelID $CHANNEL_NAME -asOrg IllinoisMSP



{"success":true,"secret":"","message":"Jim enrolled Successfully","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDcxMzM3ODYsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6ImFnZW50IiwiaWF0IjoxNTA3MDk3Nzg2fQ.qT79l2oruU8t3mn5v6esbbCuYpfgIKp8PiBbcBTePak
"}hendry@hypderledger:~$ curl -s -X POST \



curl -s -X POST \
  http://localhost:4000/channels \
  -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDcxMzM3ODYsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6ImFnZW50IiwiaWF0IjoxNTA3MDk3Nzg2fQ.qT79l2oruU8t3mn5v6esbbCuYpfgIKp8PiBbcBTePak" \
  -H "content-type: application/json" \
  -d '{
	"channelName":"mychannel",
	"channelConfigPath":"../artifacts/channel/mychannel.tx"
}'

curl -s -X POST \
  http://localhost:4000/channels/mychannel/peers \
  -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDcxMzM3ODYsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6ImFnZW50IiwiaWF0IjoxNTA3MDk3Nzg2fQ.qT79l2oruU8t3mn5v6esbbCuYpfgIKp8PiBbcBTePak" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer1","peer2"]

}'

curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDcxMzM3ODYsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6ImFnZW50IiwiaWF0IjoxNTA3MDk3Nzg2fQ.qT79l2oruU8t3mn5v6esbbCuYpfgIKp8PiBbcBTePak" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer1","peer2"],
	"chaincodeName":"mycc",
	"chaincodePath":"github.com/example_cc",
	"chaincodeVersion":"v0"
}'
curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes \
  -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDcxMzM3ODYsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6ImFnZW50IiwiaWF0IjoxNTA3MDk3Nzg2fQ.qT79l2oruU8t3mn5v6esbbCuYpfgIKp8PiBbcBTePak" \
  -H "content-type: application/json" \
  -d '{
	"chaincodeName":"mycc",
	"chaincodeVersion":"v0",
	"args":["a","100","b","200"]
}'

curl -s -X GET \
  "http://localhost:4000/channels/mychannel/chaincodes/mycc?peer=peer1&fcn=query&args=%5B%22a%22%5D" \
  -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDcxMzM3ODYsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6ImFnZW50IiwiaWF0IjoxNTA3MDk3Nzg2fQ.qT79l2oruU8t3mn5v6esbbCuYpfgIKp8PiBbcBTePak" \
  -H "content-type: application/json"
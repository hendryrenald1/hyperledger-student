
 package main
 
 import (
	 "fmt"

	 "github.com/hyperledger/fabric/core/chaincode/shim"
	 "github.com/hyperledger/fabric/protos/peer"
 )
 var logger = shim.NewLogger("sample_cc")

 type SimpleAsset struct {
 }
 
 
 func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {

	logger.Info("########### Sample Code Init Successfully ###########")
	 return shim.Success(nil)
 }
 
 
 func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	logger.Info("########### Sample Code Invoked Successfully ###########")
	 return shim.Success(nil);
 }
 

 
 // main function starts up the chaincode in the container during instantiate
 func main() {
	 if err := shim.Start(new(SimpleAsset)); err != nil {
		 fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	 }
 }
 
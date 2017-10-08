/*
 * Copyright IBM Corp All Rights Reserved
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"fmt"
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// SimpleAsset implements a simple chaincode to manage an asset
type SimpleAsset struct {
}


type University struct {
	UName string
}


type Student struct {
    Name string
    Universities []University
}

 
// Init is called during chaincode instantiation to initialize any
// data. Note that chaincode upgrade also calls this function to reset
// or to migrate data.
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	// Get the args from the transaction proposal
	args := stub.GetStringArgs()
	if len(args) != 2 {
		return shim.Error("Incorrect arguments. Expecting a key and a value")
	}

// Sample code to check if it working
	res1D := &Student{
        Name:   "Testing",
        Universities: []University{University{UName :"First University"},University{UName :"Second University"} }}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(" Modified  " , string(res1B))

	res := &Student{}
	
	json.Unmarshal([]byte(args[1]), &res)

		fmt.Println("Modified String",res)
		
		
		responseToWrite, _ := json.Marshal(res)

	// Set up any variables or assets here by calling stub.PutState()

	// We store the key and the value on the ledger
	err := stub.PutState(args[0], []byte(responseToWrite))
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
	}
	return shim.Success(nil)
}

// Invoke is called per transaction on the chaincode. Each transaction is
// either a 'get' or a 'set' on the asset created by Init function. The Set
// method may create a new asset by specifying a new key-value pair.
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// Extract the function and args from the transaction proposal
	fn, args := stub.GetFunctionAndParameters()

	var result string
	var err error
	if fn == "add" {
		result, err = add(stub, args)
	} else { // assume 'get' even if fn is nil
		result, err = get(stub, args)
	}
	if err != nil {
		return shim.Error(err.Error())
	}

	// Return the result as success payload
	return shim.Success([]byte(result))
}

// Set stores the asset (both key and value) on the ledger. If the key exists,
// it will override the value with the new one
func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key and a value")
	}

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return "", fmt.Errorf("Failed to set asset: %s", args[0])
	}
	return args[1], nil
}

// Set stores the asset (both key and value) on the ledger. If the key exists,
// it will override the value with the new one
func add(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key and a value")
	}

	//Unmarshall Student University inforamtion
	univ := &University{}
	json.Unmarshal([]byte(args[1]), &univ)

	// Get the student from the ledger

	student, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
	}
	if student == nil {
		return "", fmt.Errorf("Asset not found: %s", args[0])
	}
	// After Error check, Unmarshall to the Student data

	fmt.Println("Restrived Student In JSON format :", student )

	stud := &Student{}
	json.Unmarshal([]byte(student), &stud)

	fmt.Println("Restrived Student After Unmarshalling :", stud )

	//Append the University
	stud.Universities = append(stud.Universities,*univ) 

	fmt.Println("Restrived Student After Appending Univesity :", stud )


	updatedstudjson, _ := json.Marshal(stud)
	
	fmt.Println("Student information after Marshalling and before updating :", stud )
	

	ledgererr := stub.PutState(args[0], []byte(updatedstudjson))
	if ledgererr != nil {
		return "", fmt.Errorf("Failed to set asset: %s", args[0])
	}
	return args[1], nil
}


// Get returns the value of the specified asset key
func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
	}
	if value == nil {
		return "", fmt.Errorf("Asset not found: %s", args[0])
	}
	return string(value), nil
}

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}

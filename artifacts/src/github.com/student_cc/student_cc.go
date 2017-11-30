/*
 * Copyright IBM Corp All Rights Reserved
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// SimpleAsset implements a simple chaincode to manage an asset
type SimpleAsset struct {
}

type Student struct {
	Name         string
	id           int
	age          int
	Universities []University
}

type University struct {
	UName     string
	UAddreess string
}

var logger = shim.NewLogger("student_cc")

// Init is called during chaincode instantiation to initialize any
// data. Note that chaincode upgrade also calls this function to reset
// or to migrate data.
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	// Get the args from the transaction proposal

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

	logger.Info("Function Name returned %s", fn)

	if fn == "create" {
		result, err = create(stub, args)
	} else if fn == "add" {
		result, err = add(stub, args)
	} else if fn == "get" {
		result, err = get(stub, args)
	}
	if err != nil {
		return shim.Error(err.Error())
	}

	// Return the result as success payload
	return shim.Success([]byte(result))
}

/*
Create the student info
*/

func create(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key and a value")
	}

	fmt.Println(":: Argument from NodeJS application %s and %s :", args[0], args[1])
	fmt.Println(":: Length of the Argument %d", len(args))

	logger.Info("########### Student Init ###########")
	logger.Info(args[0])
	logger.Info(args[1])

	studentNo := args[0]
	studentInfo := args[1]

	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key and a value")
		//	return shim.Error("")
	}

	// Sample code to check if it working
	res1D := &Student{
		Name:         "Testing",
		Universities: []University{University{UName: "First University"}, University{UName: "Second University"}}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(" Modified  ", string(res1B))

	res := &Student{}

	json.Unmarshal([]byte(studentInfo), &res)

	logger.Info("Modified String %s", res)

	responseToWrite, _ := json.Marshal(res)

	// Set up any variables or assets here by calling stub.PutState()

	// We store the key and the value on the ledger
	err := stub.PutState(studentNo, []byte(responseToWrite))
	if err != nil {
		//	return shim.Error(fmt.Sprintf("Failed to create asset: %s", studentNo))
		return "", fmt.Errorf("Failed to create asset: %s", studentNo)
	} else {
		logger.Info(" Successfully written to Ledger with paramater as")
		logger.Info(studentNo)
	}
	return args[1], nil
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

	fmt.Println("Restrived Student In JSON format :", student)

	stud := &Student{}
	json.Unmarshal([]byte(student), &stud)

	fmt.Println("Restrived Student After Unmarshalling :", stud)

	//Append the University
	stud.Universities = append(stud.Universities, *univ)

	fmt.Println("Restrived Student After Appending Univesity :", stud)

	updatedstudjson, _ := json.Marshal(stud)

	fmt.Println("Student information after Marshalling and before updating :", stud)

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

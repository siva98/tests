package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Test struct {
}

func (t *Test) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success([]byte("success"))
}

func (t *Test) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	tByte, err := json.Marshal(time.Now())
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState("key", tByte)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte("success"))
}

func main() {
	if err := shim.Start(new(Test)); err != nil {
		fmt.Printf("Error: %s", err)
	}
}

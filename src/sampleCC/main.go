package main

import (
	"fmt"
	statemgr "sampleCC/stateMgr"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("sampleCCLog")

// SampleCC chaincode structure
type SampleCC struct {
	StateMgr statemgr.StateMgr
}

// Init initializes chaincode
func (t *SampleCC) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### SampleCC Init ###########")
	methodName := "Init()"
	_, args := stub.GetFunctionAndParameters()
	t.StateMgr = statemgr.StateMgr{args}

	if len(t.StateMgr.Namespaces) == 0 {
		warningMsg := fmt.Sprintf("%s - No namespaces were provided to SampleCC.", methodName)
		logger.Warning(warningMsg)
	}

	logger.Infof("Namespaces provided to SampleCC: %v", t.StateMgr.Namespaces)
	logger.Infof("- End execution -  %s\n", methodName)
	return shim.Success(nil)
}

// Invoke is the entry point for all invocations
func (t *SampleCC) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, _ := stub.GetFunctionAndParameters()
	logger.Info("########### SampleCC Invoke ###########")

	switch function {
	case "DeleteState":
		logger.Info("########### Calling DeleteState ###########")
		return t.StateMgr.DeleteState(stub)
	}

	errorMsg := fmt.Sprintf("Could not find function named '%s' in SampleCC.", function)
	logger.Errorf(errorMsg)
	return shim.Error(errorMsg)
}

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(SampleCC))
	if err != nil {
		logger.Errorf("Error starting SampleCC: %s", err)
	}
}

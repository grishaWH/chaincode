func (t *SimpleChaincode) getAllHistoryKey(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var key string
	
	key = args[0]

	historyIter, err := stub.GetHistoryForKey(key)

    	if err != nil {
        	errMsg := fmt.Sprintf("[error] cannot retrieve history for key <%s>, due to %s", key, err)
        	fmt.Println(errMsg)
        	return shim.Error(errMsg)
    	}

    	for historyIter.HasNext() {
        	modification, err := historyIter.Next()
        	if err != nil {
            		errMsg := fmt.Sprintf("[error] cannot read record modification for key %s, id <%s>, due to %s", key, err)
            		fmt.Println(errMsg)
            		return shim.Error(errMsg)
        	}
        	fmt.Println("Returning information about", string(modification.Value))
    	}


	return shim.Success(nil)
}

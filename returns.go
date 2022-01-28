package replica

// Set the return value of a specific Mock function.
func (m *mockData) SetReturnValues(fnName string, value ...interface{}) {
	m.MockReturnValues[fnName] = value
}

// Get the return values for a specific Mock function.
// Since everything will be converted to an interface{}
// when you receive this data within the mocked function,
// you should cast it to the appropriate types for your
// use case.
func (m *mockData) GetReturnValues(fnName string) []interface{} {
	return m.MockReturnValues[fnName]
}
